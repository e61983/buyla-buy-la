(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-0bc51c42"],{"1c8c":function(t,e,s){},add6:function(t,e,s){"use strict";s.r(e);var o=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("section",[s("div",{staticClass:"row"},[null!==t.current_record?s("div",{staticClass:"col-12"},[s("OrderEditor",{attrs:{record:t.current_record}})],1):t._e(),null!==t.records?s("div",{staticClass:"col-12"},t._l(t.records,(function(t,e){return s("div",{key:e},[s("Record",{attrs:{record:t}})],1)})),0):t._e()]),s("div",{staticClass:"text-left"},[s("font-awesome-icon",{staticClass:"fab fa-line",class:{"text-success":null!==t.liff_context,"text-secondary":null===t.liff_context},attrs:{icon:["fab","line"],"data-toggle":"collapse","data-target":"#context",role:"button"}}),s("div",{staticClass:"collapse multi-collapse",attrs:{id:"context"}},[s("div",{staticClass:"card card-body"},[s("span",[t._v("liff context: "+t._s(t.liff_context))]),s("span",[t._v("order group: "+t._s(t.order_group))])])])],1)])},r=[],n=(s("4160"),s("b0c0"),s("b64b"),s("159b"),function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",{staticClass:"row d-flex flex-row text-center border mt-2",staticStyle:{"border-radius":"5%"}},[s("div",{staticClass:"col-md-3 col-12 text-center"},[s("img",{staticClass:"mx-1",staticStyle:{"border-radius":"100%"},attrs:{src:t.record.user_profile.photo_url,width:"128px",height:"128px"}}),s("div",{staticClass:"h5"},[t._v(t._s(t.record.user_profile.display_name))])]),s("ValidationObserver",{staticClass:"col justify-content-center",scopedSlots:t._u([{key:"default",fn:function(e){var o=e.handleSubmit,r=e.reset;return[s("form",{on:{submit:function(e){return e.preventDefault(),o(t.submit)},reset:function(t){return t.preventDefault(),r(t)}}},[t._l(t.record.goods,(function(t,e){return s("OrderItem",{key:t.id,attrs:{good:t,index:e}})})),s("div",{staticClass:"col text-center"},[s("button",{staticClass:"btn btn-outline-success btn-block mt-2",class:{"bounce animated":t.animated},attrs:{type:"button"},on:{click:function(e){return t.create_new_good()}}},[t._v("新增一個新的品項")])]),s("div",{staticClass:"col text-center"},[s("button",{staticClass:"btn btn-outline-secondary m-1",attrs:{type:"reset"},on:{click:function(e){return t.reset_goods()}}},[t._v("重置")]),"new"===t.order_status?s("button",{staticClass:"btn btn-success m-1",attrs:{type:"submit"}},[t._v("送出")]):t._e(),"modify"===t.order_status?s("button",{staticClass:"btn btn-success m-1",attrs:{type:"submit"}},[t._v("修改")]):t._e(),"delete"===t.order_status?s("button",{staticClass:"btn btn-danger m-1",attrs:{type:"submit"}},[t._v("我不要了")]):t._e()])],2)]}}])}),s("OrderComfirm",{attrs:{id:"comfirmModal",record:t.record}})],1)}),a=[],i=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",{staticClass:"my-1"},[t.is_editor?s("div",{staticClass:"form-row border text-left",staticStyle:{"border-radius":"5px"}},[s("div",{staticClass:"col-12"},[s("div",{staticClass:"form-group"},[s("label",{attrs:{for:"item_name"+t.index}},[t._v("*品項")]),s("ValidationProvider",{attrs:{name:"品名"+(t.index+1),rules:"required"},scopedSlots:t._u([{key:"default",fn:function(e){var o=e.errors;return[s("input",{directives:[{name:"model",rawName:"v-model",value:t.good.item_name,expression:"good.item_name"}],staticClass:"form-control",attrs:{type:"text",id:"item_name"+t.good.id,placeholder:"請輸入品名"+(t.index+1)},domProps:{value:t.good.item_name},on:{input:function(e){e.target.composing||t.$set(t.good,"item_name",e.target.value)}}}),s("span",{staticClass:"text-danger"},[t._v(t._s(o[0]))])]}}],null,!1,3845636014)})],1)]),s("div",{staticClass:"col-6 col-md-3"},[s("div",{staticClass:"form-group"},[s("label",{attrs:{for:"sweeting_level"+t.index}},[t._v("甜度")]),s("select",{directives:[{name:"model",rawName:"v-model",value:t.good.sweetness_level,expression:"good.sweetness_level"}],staticClass:"form-control",attrs:{type:"text",id:"sweeting_level"+t.index},on:{change:function(e){var s=Array.prototype.filter.call(e.target.options,(function(t){return t.selected})).map((function(t){var e="_value"in t?t._value:t.value;return e}));t.$set(t.good,"sweetness_level",e.target.multiple?s:s[0])}}},[s("option",{attrs:{value:"正常",selected:""}},[t._v("正常")]),s("option",{attrs:{value:"少糖"}},[t._v("少糖")]),s("option",{attrs:{value:"半糖"}},[t._v("半糖")]),s("option",{attrs:{value:"微糖"}},[t._v("微糖")]),s("option",{attrs:{value:"無糖"}},[t._v("無糖")])])])]),s("div",{staticClass:"col-6 col-md-3"},[s("div",{staticClass:"form-group"},[s("label",{attrs:{for:"amount_of_icd"+t.index}},[t._v("冰量")]),s("select",{directives:[{name:"model",rawName:"v-model",value:t.good.amount_of_ice,expression:"good.amount_of_ice"}],staticClass:"form-control",attrs:{type:"text",id:"amount_of_icd"+t.index},on:{change:function(e){var s=Array.prototype.filter.call(e.target.options,(function(t){return t.selected})).map((function(t){var e="_value"in t?t._value:t.value;return e}));t.$set(t.good,"amount_of_ice",e.target.multiple?s:s[0])}}},[s("option",{attrs:{value:"正常",selected:""}},[t._v("正常")]),s("option",{attrs:{value:"少冰"}},[t._v("少冰")]),s("option",{attrs:{value:"去冰"}},[t._v("去冰")]),s("option",{attrs:{value:"微冰"}},[t._v("微冰")]),s("option",{attrs:{value:"去冰"}},[t._v("去冰")]),s("option",{attrs:{value:"完全去冰"}},[t._v("完全去冰")]),s("option",{attrs:{value:"熱的"}},[t._v("熱的")])])])]),s("div",{staticClass:"col-6 col-md-3"},[s("div",{staticClass:"form-group"},[s("label",{attrs:{for:"size"+t.index}},[t._v("大小")]),s("select",{directives:[{name:"model",rawName:"v-model",value:t.good.size,expression:"good.size"}],staticClass:"form-control",attrs:{type:"text",id:"size"+t.index},on:{change:function(e){var s=Array.prototype.filter.call(e.target.options,(function(t){return t.selected})).map((function(t){var e="_value"in t?t._value:t.value;return e}));t.$set(t.good,"size",e.target.multiple?s:s[0])}}},[s("option",{attrs:{value:"L",selected:""}},[t._v("L")]),s("option",{attrs:{value:"M"}},[t._v("M")]),s("option",{attrs:{value:"S"}},[t._v("S")]),s("option",{attrs:{value:"XL"}},[t._v("XL")])])])]),s("div",{staticClass:"col-6 col-md-3"},[s("div",{staticClass:"form-group"},[s("label",{attrs:{for:"number"+t.index}},[t._v("數量")]),s("input",{directives:[{name:"model",rawName:"v-model",value:t.good.number,expression:"good.number"}],staticClass:"form-control",attrs:{type:"number",id:"number"+t.index},domProps:{value:t.good.number},on:{input:function(e){e.target.composing||t.$set(t.good,"number",e.target.value)}}})])]),s("div",{staticClass:"col-12"},[s("div",{staticClass:"form-group mb-0"},[s("label",{attrs:{for:"comment"+t.index}},[t._v("備註")]),s("input",{directives:[{name:"model",rawName:"v-model",value:t.good.comment,expression:"good.comment"}],staticClass:"form-control",attrs:{type:"text",id:"comment"+t.index,placeholder:"(Optional) 想附註的事"},domProps:{value:t.good.comment},on:{input:function(e){e.target.composing||t.$set(t.good,"comment",e.target.value)}}})])]),s("div",{staticClass:"col-md m-2 text-center"},[s("button",{staticClass:"btn btn-sm btn-outline-success btn-block",attrs:{type:"button"},on:{click:function(e){return t.triggor_editor()}}},[t._v("確認")]),s("button",{staticClass:"btn btn-sm btn-outline-danger btn-block",attrs:{type:"button"},on:{click:function(e){return t.destory_good(t.good)}}},[t._v("移除")])])]):s("div",{staticClass:"row align-items-center flex-column"},[s("div",{staticClass:"col-12 text-left"},[s("div",[t._v(" "+t._s(t.good.item_name)+" "+t._s(t.good.sweetness_level)+" "+t._s(t.good.amount_of_ice)+" "+t._s(t.good.size)+" * "+t._s(t.good.number)+" "),s("a",{staticClass:"badge badge-success mr-0",attrs:{href:"#"},on:{click:function(e){return t.triggor_editor()}}},[t._v("修改")]),s("a",{staticClass:"badge badge-danger mr-0",attrs:{href:"#"},on:{click:function(e){return t.destory_good(t.good)}}},[t._v("移除")])])])])])},c=[],l={name:"OrderItem",props:["good","index"],data:function(){return{is_editor:!0}},created:function(){""===this.good.item_name?this.is_editor=!0:this.is_editor=!1},methods:{triggor_editor:function(){this.is_editor=!this.is_editor},destory_good:function(t){this.$store.dispatch("remove_good",t.id)}}},d=l,u=s("2877"),m=Object(u["a"])(d,i,c,!1,null,null,null),_=m.exports,f=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",{staticClass:"modal fade",attrs:{id:t.id,tabindex:"-1",role:"dialog","aria-labelledby":"exampleModalLabel","aria-hidden":"true"}},[s("form",{on:{submit:function(t){t.preventDefault()}}},[s("div",{staticClass:"modal-dialog",attrs:{role:"document"}},[s("div",{staticClass:"modal-content"},[t._m(0),"delete"!==t.order_status?s("div",{staticClass:"modal-body"},[s("div",{staticClass:"card"},[s("div",{staticClass:"card-body"},[s("div",{staticClass:"row align-items-center flex-column"},[s("div",{staticClass:"col text-left"},t._l(t.record.goods,(function(e,o){return s("ul",{key:o,staticClass:"mb-0"},[s("li",[t._v(" "+t._s(e.item_name)+" "+t._s(e.sweetness_level)+" "+t._s(e.amount_of_ice)+" "+t._s(e.size)+" * "+t._s(e.number)+" ("+t._s(e.comment)+") ")])])})),0)])])])]):t._e(),s("div",{staticClass:"modal-footer"},[s("button",{staticClass:"btn btn-outline-danger",attrs:{type:"button","data-dismiss":"modal"}},[t._v("再想一下")]),"new"===t.order_status?s("button",{staticClass:"btn btn-success mx-2",attrs:{type:"button"},on:{click:function(e){return t.submit_order()}}},[t._v("確認加入揪團")]):t._e(),"modify"===t.order_status?s("button",{staticClass:"btn btn-success mx-2",attrs:{type:"button"},on:{click:function(e){return t.modify_order()}}},[t._v("修改")]):t._e(),"delete"===t.order_status?s("button",{staticClass:"btn btn-danger mx-2",attrs:{type:"button"},on:{click:function(e){return t.delete_order()}}},[t._v("我不要了")]):t._e()])])])])])},v=[function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",{staticClass:"modal-header"},[s("h5",{staticClass:"modal-title",attrs:{id:"exampleModalLabel"}},[t._v("再一次確認")])])}],p=(s("d3b7"),s("25f0"),s("3ca3"),s("ddb0"),s("2b3d"),s("1157")),g=s.n(p),b={name:"OrderComfirm",props:{id:{required:!0,type:String},record:{required:!0,type:Object}},computed:{order_status:function(){var t=this.$store.state.order_status,e=this.record.goods.length;return"exist"===t?e>0?"modify":"delete":"new"}},methods:{delete_order:function(){var t=this,e=this,s="https://only-test-0001.herokuapp.com/api/v1/",o=new URL(String(e.$store.state.gid)+"/order/"+String(e.$store.state.uid),s).toString();console.log(o),e.$http.delete(o).then((function(t){console.log("Result",t),e.$store.dispatch("get_records"),g()("#"+e.id).modal("hide")})).then((function(){t.$store.dispatch("add_message",{status:"success",message:"已經成功取消了喔"})})).then((function(){null!==e.$store.state.liff_context&&e.liff_sent_message("[我不要了]")})).catch((function(e){t.$store.dispatch("add_message",{status:"danger",message:"Ops~ 好像出了點問題"}),console.log(e)}))},modify_order:function(){var t=this,e=this,s="https://only-test-0001.herokuapp.com/api/v1/",o=new URL(String(e.$store.state.gid)+"/order/"+String(e.$store.state.uid),s).toString();console.log(o),e.$http.delete(o).then((function(t){console.log("Result",t),e.$store.dispatch("set_order_status","empty"),g()("#"+e.id).modal("hide")})).then((function(){var s="https://only-test-0001.herokuapp.com/api/v1/",o=new URL(String(e.$store.state.gid)+"/order/"+String(e.$store.state.uid),s).toString();e.$http.post(o,e.record).then((function(t){console.log("Result",t),g()("#"+e.id).modal("hide")})).then((function(){e.$store.dispatch("get_records"),t.$store.dispatch("add_message",{status:"success",message:"已經成功修改囉"})})).then((function(){if(null!==e.$store.state.liff_context){var s="",o=t;o.record.goods.forEach((function(t){s+="\n •"+t.item_name+" "+t.sweetness_level+" "+t.amount_of_ice+" "+t.size+"*"+t.number+"("+t.comment+")"})),o.liff_sent_message("[我要]"+s)}})).catch((function(e){t.$store.dispatch("add_message",{status:"danger",message:"Ops~ 好像出了點問題"}),console.log(e)}))}))},submit_order:function(){var t=this,e=this,s="https://only-test-0001.herokuapp.com/api/v1/",o=new URL(String(e.$store.state.gid)+"/order/"+String(e.$store.state.uid),s).toString();console.log(o),e.$http.post(o,e.record).then((function(t){console.log("Result",t),e.$store.dispatch("get_records"),g()("#"+e.id).modal("hide")})).then((function(){t.$store.dispatch("add_message",{status:"success",message:"已經成功 + 1 囉"})})).then((function(){if(null!==e.$store.state.liff_context){var s="",o=t;o.record.goods.forEach((function(t){s+="\n •"+t.item_name+" "+t.sweetness_level+" "+t.amount_of_ice+" "+t.size+"*"+t.number+"("+t.comment+")"})),o.liff_sent_message("[我要]"+s)}})).catch((function(e){t.$store.dispatch("add_message",{status:"danger",message:"Ops~ 好像出了點問題"}),console.log(e)}))},liff_sent_message:function(t){var e=this;e.$liff.sendMessages([{type:"text",text:t}]).then((function(){e.$liff.closeWindow()})).catch((function(t){console.log("error",t)}))}}},h=b,x=Object(u["a"])(h,f,v,!1,null,null,null),C=x.exports,y={name:"OrderEditor",components:{OrderItem:_,OrderComfirm:C},props:["record"],data:function(){return{animated:!1}},computed:{order_status:function(){var t=this.$store.state.order_status,e=this.record.goods.length;return"exist"===t?e>0?"modify":"delete":"new"}},methods:{create_new_good:function(){this.$store.dispatch("add_good",{item_name:"",sweetness_level:"正常",amount_of_ice:"正常",size:"L",number:"1",comment:""})},reset_goods:function(){this.$store.dispatch("get_records")},submit:function(){var t=this,e=t.$store.state.order_status;"new"===e?t.record.goods.length<=0&&(t.animated=!0,setTimeout((function(){t.animated=!1}),1e3)):(console.log("show check madol"),g()("#comfirmModal").modal("show"))}}},$=y,w=Object(u["a"])($,n,a,!1,null,"40b75384",null),k=w.exports,S=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",{staticClass:"row d-flex flex-row text-center border mt-2",staticStyle:{"border-radius":"5%"}},[s("div",{staticClass:"col-md-3 col-12 d-flex flex-row align-items-center"},[s("img",{staticClass:"mr-3",staticStyle:{"border-radius":"100%"},attrs:{src:t.record.user_profile.photo_url,width:"48px",height:"48px"}}),s("div",{staticClass:"h5 mt-2"},[t._v(t._s(t.record.user_profile.display_name))])]),s("div",{staticClass:"col-md-9 col-12"},[s("div",{staticClass:"row align-items-center flex-column"},[s("div",{staticClass:"col text-left"},t._l(t.record.goods,(function(e,o){return s("ul",{key:o,staticClass:"mb-0"},[s("li",[s("div",[t._v(" "+t._s(e.item_name)+" "+t._s(e.sweetness_level)+" "+t._s(e.amount_of_ice)+" "+t._s(e.size)+" * "+t._s(e.number)+" ")]),""!==e.comment?s("div",{staticClass:"border px-2",staticStyle:{"border-radius":"5%"}},[t._v(t._s(e.comment))]):t._e()])])})),0),s("div",{staticClass:"col my-2 text-center"},[s("button",{staticClass:"btn btn-outline-primary btn-sm-block",attrs:{type:"button"},on:{click:function(e){return t.same_as(t.record.goods)}}},[t._v("跟"+t._s(t.record.user_profile.display_name)+"一樣")])])])])])},O=[],E={name:"Record",props:["record"],computed:{show_records:function(){return this.record}},methods:{same_as:function(t){this.$store.dispatch("goods_same_as",t)}}},z=E,L=Object(u["a"])(z,S,O,!1,null,null,null),M=L.exports,R={name:"MainMenu",components:{OrderEditor:k,Record:M},props:{title:String},data:function(){return{is_open_editor:!1}},created:function(){console.log(this.$options.name,"created")},beforeMount:function(){console.log(this.$options.name,"beforeMount")},mounted:function(){console.log(this.$options.name,"mounted"),this.$store.dispatch("set_is_loading",!0),this.$store.dispatch("set_is_loading",!1)},computed:{records:function(){var t=this,e=t.$store.state.records;if(null===e)return null;var s=[];return Object.keys(e).forEach((function(o){o!==t.$store.state.uid&&s.push(e[o])})),s},liff_context:function(){return this.$store.state.liff_context},order_group:function(){return this.$store.state.gid},current_record:function(){var t=this;return{user_profile:t.$store.state.user_profile,goods:t.$store.state.current_goods}}}},j=R,N=(s("e436"),Object(u["a"])(j,o,r,!1,null,"44bc0fb6",null));e["default"]=N.exports},e436:function(t,e,s){"use strict";var o=s("1c8c"),r=s.n(o);r.a}}]);
//# sourceMappingURL=chunk-0bc51c42.7e80f605.js.map