"use strict";(self.webpackChunkstacklok=self.webpackChunkstacklok||[]).push([[3594],{3905:(e,r,t)=>{t.d(r,{Zo:()=>c,kt:()=>f});var n=t(67294);function i(e,r,t){return r in e?Object.defineProperty(e,r,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[r]=t,e}function o(e,r){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);r&&(n=n.filter((function(r){return Object.getOwnPropertyDescriptor(e,r).enumerable}))),t.push.apply(t,n)}return t}function l(e){for(var r=1;r<arguments.length;r++){var t=null!=arguments[r]?arguments[r]:{};r%2?o(Object(t),!0).forEach((function(r){i(e,r,t[r])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):o(Object(t)).forEach((function(r){Object.defineProperty(e,r,Object.getOwnPropertyDescriptor(t,r))}))}return e}function a(e,r){if(null==e)return{};var t,n,i=function(e,r){if(null==e)return{};var t,n,i={},o=Object.keys(e);for(n=0;n<o.length;n++)t=o[n],r.indexOf(t)>=0||(i[t]=e[t]);return i}(e,r);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(n=0;n<o.length;n++)t=o[n],r.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(i[t]=e[t])}return i}var p=n.createContext({}),s=function(e){var r=n.useContext(p),t=r;return e&&(t="function"==typeof e?e(r):l(l({},r),e)),t},c=function(e){var r=s(e.components);return n.createElement(p.Provider,{value:r},e.children)},d="mdxType",u={inlineCode:"code",wrapper:function(e){var r=e.children;return n.createElement(n.Fragment,{},r)}},m=n.forwardRef((function(e,r){var t=e.components,i=e.mdxType,o=e.originalType,p=e.parentName,c=a(e,["components","mdxType","originalType","parentName"]),d=s(t),m=i,f=d["".concat(p,".").concat(m)]||d[m]||u[m]||o;return t?n.createElement(f,l(l({ref:r},c),{},{components:t})):n.createElement(f,l({ref:r},c))}));function f(e,r){var t=arguments,i=r&&r.mdxType;if("string"==typeof e||i){var o=t.length,l=new Array(o);l[0]=m;var a={};for(var p in r)hasOwnProperty.call(r,p)&&(a[p]=r[p]);a.originalType=e,a[d]="string"==typeof e?e:i,l[1]=a;for(var s=2;s<o;s++)l[s]=t[s];return n.createElement.apply(null,l)}return n.createElement.apply(null,t)}m.displayName="MDXCreateElement"},3957:(e,r,t)=>{t.r(r),t.d(r,{assets:()=>p,contentTitle:()=>l,default:()=>u,frontMatter:()=>o,metadata:()=>a,toc:()=>s});var n=t(87462),i=(t(67294),t(3905));const o={title:"minder provider list"},l=void 0,a={unversionedId:"ref/cli/minder_provider_list",id:"ref/cli/minder_provider_list",title:"minder provider list",description:"minder provider list",source:"@site/docs/ref/cli/minder_provider_list.md",sourceDirName:"ref/cli",slug:"/ref/cli/minder_provider_list",permalink:"/ref/cli/minder_provider_list",draft:!1,tags:[],version:"current",frontMatter:{title:"minder provider list"},sidebar:"minder",previous:{title:"minder provider enroll",permalink:"/ref/cli/minder_provider_enroll"},next:{title:"minder quickstart",permalink:"/ref/cli/minder_quickstart"}},p={},s=[{value:"minder provider list",id:"minder-provider-list",level:2},{value:"Synopsis",id:"synopsis",level:3},{value:"Options",id:"options",level:3},{value:"Options inherited from parent commands",id:"options-inherited-from-parent-commands",level:3},{value:"SEE ALSO",id:"see-also",level:3}],c={toc:s},d="wrapper";function u(e){let{components:r,...t}=e;return(0,i.kt)(d,(0,n.Z)({},c,t,{components:r,mdxType:"MDXLayout"}),(0,i.kt)("h2",{id:"minder-provider-list"},"minder provider list"),(0,i.kt)("p",null,"List the providers available in a specific project"),(0,i.kt)("h3",{id:"synopsis"},"Synopsis"),(0,i.kt)("p",null,"The minder provider list command lists the providers available in a specific project."),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre"},"minder provider list [flags]\n")),(0,i.kt)("h3",{id:"options"},"Options"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre"},'  -h, --help            help for list\n  -o, --output string   Output format (one of json,yaml,table) (default "table")\n')),(0,i.kt)("h3",{id:"options-inherited-from-parent-commands"},"Options inherited from parent commands"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre"},'      --config string            Config file (default is $PWD/config.yaml)\n      --grpc-host string         Server host (default "api.stacklok.com")\n      --grpc-insecure            Allow establishing insecure connections\n      --grpc-port int            Server port (default 443)\n      --identity-client string   Identity server client ID (default "minder-cli")\n      --identity-url string      Identity server issuer URL (default "https://auth.stacklok.com")\n  -j, --project string           ID of the project\n  -p, --provider string          Name of the provider, i.e. github (default "github")\n')),(0,i.kt)("h3",{id:"see-also"},"SEE ALSO"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"/ref/cli/minder_provider"},"minder provider"),"\t - Manage providers within a minder control plane")))}u.isMDXComponent=!0}}]);