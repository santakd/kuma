package istio

import (
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/kumahq/kuma/pkg/transparentproxy/istio/tools/istio-iptables/pkg/constants"

	uninstall "github.com/kumahq/kuma/pkg/transparentproxy/istio/tools/istio-clean-iptables/pkg/cmd"
	install "github.com/kumahq/kuma/pkg/transparentproxy/istio/tools/istio-iptables/pkg/cmd"

	"github.com/kumahq/kuma/pkg/transparentproxy/config"
)

type IstioTransparentProxy struct {
	// output redirection
	stdout *os.File
	stderr *os.File
	reader *os.File
}

func NewIstioTransparentProxy() *IstioTransparentProxy {
	return &IstioTransparentProxy{}
}

func (tp *IstioTransparentProxy) Setup(cfg *config.TransparentProxyConfig) (string, error) {
	viper.Set(constants.EnvoyPort, cfg.RedirectPortOutBound)
	viper.Set(constants.InboundCapturePort, cfg.RedirectPortInBound)
	viper.Set(constants.InboundCapturePortV6, cfg.RedirectPortInBoundV6)
	viper.Set(constants.ProxyUID, cfg.UID)
	viper.Set(constants.ProxyGID, cfg.GID)
	viper.Set(constants.InboundInterceptionMode, "REDIRECT")
	if cfg.RedirectInBound {
		viper.Set(constants.InboundPorts, "*")
	}
	viper.Set(constants.LocalExcludePorts, cfg.ExcludeInboundPorts)
	viper.Set(constants.ServiceCidr, "*")
	viper.Set(constants.LocalOutboundPortsExclude, cfg.ExcludeOutboundPorts)
	viper.Set(constants.DryRun, cfg.DryRun)
	viper.Set(constants.SkipRuleApply, false)
	viper.Set(constants.RunValidation, false)
	viper.Set(constants.RedirectDNS, cfg.RedirectDNS)
	viper.Set(constants.RedirectAllDNSTraffic, cfg.RedirectAllDNSTraffic)
	viper.Set(constants.AgentDNSListenerPort, cfg.AgentDNSListenerPort)
	viper.Set(constants.DNSUpstreamTargetChain, cfg.DNSUpstreamTargetChain)

	if !cfg.Verbose {
		tp.redirectStdOutStdErr()
		defer func() {
			tp.restoreStdOutStderr()
		}()
	}

	savedArgs := os.Args[1:]
	os.Args = os.Args[:1]
	defer func() {
		os.Args = append(os.Args, savedArgs...)
	}()

	if err := install.GetCommand().Execute(); err != nil {
		return tp.getStdOutStdErr(), errors.Wrapf(err, "setting istio")
	}

	return tp.getStdOutStdErr(), nil
}

func (tp *IstioTransparentProxy) Cleanup(dryRun, verbose bool) (string, error) {

	viper.Set(constants.DryRun, dryRun)
	viper.Set(constants.DNSUpstreamTargetChain, "")

	if !verbose {
		tp.redirectStdOutStdErr()
		defer func() {
			tp.restoreStdOutStderr()
		}()
	}

	savedArgs := os.Args[1:]
	os.Args = os.Args[:1]
	defer func() {
		os.Args = append(os.Args, savedArgs...)
	}()

	if err := uninstall.GetCommand().Execute(); err != nil {
		return tp.getStdOutStdErr(), errors.Wrapf(err, "setting istio")
	}

	return tp.getStdOutStdErr(), nil
}

func (tp *IstioTransparentProxy) redirectStdOutStdErr() {
	reader, writer, err := os.Pipe()

	if err != nil {
		panic(err)
	}
	tp.stdout = os.Stdout
	tp.stderr = os.Stderr
	tp.reader = reader

	os.Stdout = writer
	os.Stderr = writer
}

func (tp *IstioTransparentProxy) getStdOutStdErr() string {
	data := make([]byte, 1*1024*1024)

	_, _ = tp.reader.Read(data)

	return string(data)
}

func (tp *IstioTransparentProxy) restoreStdOutStderr() {
	os.Stdout = tp.stdout
	os.Stderr = tp.stderr
}
