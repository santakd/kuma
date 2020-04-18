(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-236cabf6"],{"362e":function(t,e,a){"use strict";a.r(e);var i=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"all-meshes"},[a("FrameSkeleton",[a("DataOverview",{attrs:{"page-size":6,"has-error":t.hasError,"is-loading":t.isLoading,"is-empty":t.isEmpty,"empty-state":t.empty_state,"display-data-table":!0,"table-data":t.tableData,"table-data-is-empty":t.tableDataIsEmpty,"table-data-function-text":"View","table-data-row":"name"},on:{tableAction:t.tableAction,reloadData:t.bootstrap}},[a("template",{slot:"additionalControls"},[a("KButton",{staticClass:"add-mesh-button",attrs:{appearance:"primary",size:"small",to:{path:"/wizard/mesh"}}},[t._v("\n          Create Mesh\n        ")])],1)],2),a("Tabs",{attrs:{"has-error":t.hasError,"is-loading":t.isLoading,"is-empty":t.isEmpty,tabs:t.tabs,"tab-group-title":t.tabGroupTitle,"initial-tab-override":"overview"}},[a("template",{slot:"overview"},[a("LabelList",{attrs:{"has-error":t.entityHasError,"is-loading":t.entityIsLoading,"is-empty":t.entityIsEmpty,items:t.entity}})],1),a("template",{slot:"yaml"},[a("YamlView",{attrs:{title:t.entityOverviewTitle,"has-error":t.entityHasError,"is-loading":t.entityIsLoading,"is-empty":t.entityIsEmpty,content:t.rawEntity}})],1)],2)],1)],1)},n=[],s=a("75fc"),r=(a("7f7f"),a("d7c2")),o=a("8218"),l=a("1d10"),y=a("2778"),c=a("251b"),m=a("ff9d"),u=a("0ada"),h={name:"Meshes",metaInfo:{title:"Meshes"},components:{FrameSkeleton:l["a"],DataOverview:y["a"],Tabs:c["a"],YamlView:m["a"],LabelList:u["a"]},mixins:[o["a"]],data:function(){return{isLoading:!0,isEmpty:!1,hasError:!1,entityIsLoading:!0,entityIsEmpty:!1,entityHasError:!1,tableDataIsEmpty:!1,empty_state:{title:"No Data",message:"There are no Meshes present."},tableData:{headers:[{key:"actions",hideLabel:!0},{label:"Name",key:"name"},{label:"Type",key:"type"}],data:[]},tabs:[{hash:"#overview",title:"Overview"},{hash:"#yaml",title:"YAML"}],entity:null,rawEntity:null,firstEntity:null}},computed:{tabGroupTitle:function(){var t=this.entity;return t?"Meshes: ".concat(t.name):null},entityOverviewTitle:function(){var t=this.entity;return t?"Entity Overview for ".concat(t.name):null}},watch:{$route:function(t,e){this.bootstrap()}},beforeMount:function(){this.bootstrap()},methods:{tableAction:function(t){var e=t;this.$store.dispatch("updateSelectedTab",this.tabs[0].hash),this.$store.dispatch("updateSelectedTableRow",e.name),this.getEntity(e)},bootstrap:function(){var t=this;this.isLoading=!0,this.isEmpty=!1;var e=this.$route.params.mesh,a="all"!==e&&e?this.$api.getMesh(e):this.$api.getAllMeshes(),i=function(){return a.then((function(a){var i=function(){if("all"===e)return a.items;var t={items:[]};return t.items.push(a),t.items},n=i();n.length>0?("all"===e&&t.sortEntities(n),t.firstEntity=n[0].name,t.getEntity(n[0]),t.$store.dispatch("updateSelectedTableRow",t.firstEntity),t.tableData.data=Object(s["a"])(n),t.tableDataIsEmpty=!1):(t.tableData.data=[],t.tableDataIsEmpty=!0,t.getEntity(null))})).catch((function(e){t.hasError=!0,console.error(e)})).finally((function(){setTimeout((function(){t.isLoading=!1}),"500")}))};i()},getEntity:function(t){var e=this;if(this.entityIsLoading=!0,this.entityIsEmpty=!1,t&&null!==t)return this.$api.getMesh(t.name).then((function(t){if(t){var a=["type","name"];e.entity=Object(r["a"])(t,a),e.rawEntity=t}else e.entity=null,e.entityIsEmpty=!0})).catch((function(t){e.entityHasError=!0,console.error(t)})).finally((function(){setTimeout((function(){e.entityIsLoading=!1}),"500")}));setTimeout((function(){e.entityIsEmpty=!0,e.entityIsLoading=!1}),"500")}}},p=h,d=(a("547f"),a("2877")),b=Object(d["a"])(p,i,n,!1,null,"302567e4",null);e["default"]=b.exports},"547f":function(t,e,a){"use strict";var i=a("9839"),n=a.n(i);n.a},9839:function(t,e,a){}}]);