(function(t){function e(e){for(var n,o,s=e[0],u=e[1],l=e[2],v=0,p=[];v<s.length;v++)o=s[v],Object.prototype.hasOwnProperty.call(a,o)&&a[o]&&p.push(a[o][0]),a[o]=0;for(n in u)Object.prototype.hasOwnProperty.call(u,n)&&(t[n]=u[n]);c&&c(e);while(p.length)p.shift()();return i.push.apply(i,l||[]),r()}function r(){for(var t,e=0;e<i.length;e++){for(var r=i[e],n=!0,s=1;s<r.length;s++){var u=r[s];0!==a[u]&&(n=!1)}n&&(i.splice(e--,1),t=o(o.s=r[0]))}return t}var n={},a={app:0},i=[];function o(e){if(n[e])return n[e].exports;var r=n[e]={i:e,l:!1,exports:{}};return t[e].call(r.exports,r,r.exports,o),r.l=!0,r.exports}o.m=t,o.c=n,o.d=function(t,e,r){o.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:r})},o.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},o.t=function(t,e){if(1&e&&(t=o(t)),8&e)return t;if(4&e&&"object"===typeof t&&t&&t.__esModule)return t;var r=Object.create(null);if(o.r(r),Object.defineProperty(r,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var n in t)o.d(r,n,function(e){return t[e]}.bind(null,n));return r},o.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return o.d(e,"a",e),e},o.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},o.p="/";var s=window["webpackJsonp"]=window["webpackJsonp"]||[],u=s.push.bind(s);s.push=e,s=s.slice();for(var l=0;l<s.length;l++)e(s[l]);var c=u;i.push([0,"chunk-vendors"]),r()})({0:function(t,e,r){t.exports=r("6868")},"3a0e":function(t,e,r){"use strict";var n=r("c394"),a=r.n(n);a.a},c394:function(t,e,r){},e76a:function(t,e,r){"use strict";var n=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",{staticClass:"bg"},[t._l(t.items,(function(e,n){return r("form",{key:n,attrs:{id:"name"}},[t._m(0,!0),r("br"),r("button",{attrs:{type:"button",id:"add"},on:{click:t.new_item}},[t._v("+新增品項")]),r("br"),r("br"),r("br")])})),r("input",{staticClass:"submit_bt",attrs:{type:"submit",value:"送出",onclick:"show_alert()"}})],2)},a=[function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",{staticClass:"form1",attrs:{id:"form1"}},[r("div",{staticClass:"div1"},[r("h2",[t._v("品名")]),r("input",{attrs:{id:"drink",type:"text",name:"drinkname"}})]),r("div",{staticClass:"div2"},[r("h2",[t._v("甜度")]),r("select",{attrs:{id:"sugar_le",name:"sugar"}},[r("option",{attrs:{value:"100%sugar"}},[t._v("正常糖 100%")]),r("option",{attrs:{value:"70%sugar"}},[t._v("少糖 70%")]),r("option",{attrs:{value:"50%sugar"}},[t._v("半糖 50%")]),r("option",{attrs:{value:"30%sugar"}},[t._v("微糖 30%")]),r("option",{attrs:{value:"0%sugar"}},[t._v("無糖 0%")])])]),r("div",{staticClass:"div3"},[r("h2",[t._v("溫度")]),r("select",{attrs:{id:"temperature_le",name:"temperature"}},[r("option",{attrs:{value:"Hot"}},[t._v("熱 Hot")]),r("option",{attrs:{value:"Warm"}},[t._v("溫 Warm")]),r("option",{attrs:{value:"Without Ice"}},[t._v("去冰 Without Ice")]),r("option",{attrs:{value:"Half Ice"}},[t._v("微冰 Half Ice")]),r("option",{attrs:{value:"Less Ice"}},[t._v("少冰 Less Ice")]),r("option",{attrs:{value:"Regular Ice"}},[t._v("正常冰 Regular Ice")])])]),r("div",{staticClass:"div4"},[r("h2",[t._v("數量")]),r("input",{attrs:{type:"text",name:"quantity"}})]),r("div",{staticClass:"div5"},[r("h2",[t._v("加配料")]),r("input",{attrs:{id:"add_on",type:"text",name:"drinkaddon"}})]),r("div",{staticClass:"div6"},[r("h2",[t._v("附註")]),r("input",{attrs:{id:"other",type:"text",name:"drinkother"}})]),r("br"),r("hr")])}],i={data:function(){return{items:[{}]}},methods:{new_item:function(){this.items.push({})}}},o=i,s=(r("3a0e"),r("95c6")),u=Object(s["a"])(o,n,a,!1,null,"48cd4196",null);e["a"]=u.exports}});
//# sourceMappingURL=app.e14fb5bc.js.map