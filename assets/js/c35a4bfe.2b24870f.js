"use strict";(self.webpackChunkminder_docs=self.webpackChunkminder_docs||[]).push([[1854],{11252:(e,r,t)=>{t.r(r),t.d(r,{assets:()=>c,contentTitle:()=>o,default:()=>p,frontMatter:()=>s,metadata:()=>a,toc:()=>l});var n=t(74848),i=t(28453);const s={title:"Creating your first profile",sidebar_position:60},o="Creating your first profile",a={id:"getting_started/first_profile",title:"Creating your first profile",description:"Once you have registered repositories, you can create a profile that specifies the common, consistent configuration that you expect your your repositories to comply with.",source:"@site/docs/getting_started/first_profile.md",sourceDirName:"getting_started",slug:"/getting_started/first_profile",permalink:"/getting_started/first_profile",draft:!1,unlisted:!1,tags:[],version:"current",sidebarPosition:60,frontMatter:{title:"Creating your first profile",sidebar_position:60},sidebar:"minder",previous:{title:"Registering repositories",permalink:"/getting_started/register_repos"},next:{title:"Viewing profile status",permalink:"/getting_started/viewing_status"}},c={},l=[{value:"Prerequisites",id:"prerequisites",level:2},{value:"Creating a profile",id:"creating-a-profile",level:2},{value:"Adding rule types",id:"adding-rule-types",level:3},{value:"Creating a profile",id:"creating-a-profile-1",level:3}];function d(e){const r={a:"a",code:"code",h1:"h1",h2:"h2",h3:"h3",header:"header",p:"p",pre:"pre",...(0,i.R)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(r.header,{children:(0,n.jsx)(r.h1,{id:"creating-your-first-profile",children:"Creating your first profile"})}),"\n",(0,n.jsx)(r.p,{children:"Once you have registered repositories, you can create a profile that specifies the common, consistent configuration that you expect your your repositories to comply with."}),"\n",(0,n.jsx)(r.h2,{id:"prerequisites",children:"Prerequisites"}),"\n",(0,n.jsxs)(r.p,{children:["Before you can create a profile, you should ",(0,n.jsx)(r.a,{href:"register_repos",children:"register repositories"}),"."]}),"\n",(0,n.jsx)(r.h2,{id:"creating-a-profile",children:"Creating a profile"}),"\n",(0,n.jsx)(r.p,{children:"A profile is composed of a set of one or more rule types, each of which scan for an individual piece of configuration within a repository."}),"\n",(0,n.jsx)(r.p,{children:"For example, you may have a profile that describes your organization's security best practices, and this profile may have two rule types: GitHub secret scanning should be enabled, and secret push protection should be enabled."}),"\n",(0,n.jsx)(r.p,{children:"In this example, Minder will scan the repositories that you've registered and identify any repositories that do not have secret scanning enabled, and those that do not have secret push protection enabled."}),"\n",(0,n.jsx)(r.h3,{id:"adding-rule-types",children:"Adding rule types"}),"\n",(0,n.jsxs)(r.p,{children:["In Minder, the rule type configuration is completely flexible, and you can author them yourself. Because of this, your Minder organization does not have any rule types configured when you create it. You need to configure some, and you can upload some of Minder's already created rules from the ",(0,n.jsx)(r.a,{href:"https://github.com/mindersec/minder-rules-and-profiles",children:"minder-rules-and-profiles repository"}),"."]}),"\n",(0,n.jsxs)(r.p,{children:["For example, to add a rule type that ensures that secret scanning is enabled, and one that ensures that secret push protection is enabled, you can download the ",(0,n.jsx)(r.a,{href:"https://github.com/mindersec/minder-rules-and-profiles/blob/main/rule-types/github/secret_scanning.yaml",children:"secret_scanning.yaml"})," and ",(0,n.jsx)(r.a,{href:"https://github.com/mindersec/minder-rules-and-profiles/blob/main/rule-types/github/secret_push_protection.yaml",children:"secret_push_protection.yaml"})," rule types."]}),"\n",(0,n.jsx)(r.pre,{children:(0,n.jsx)(r.code,{className:"language-bash",children:"curl -LO https://raw.githubusercontent.com/mindersec/minder-rules-and-profiles/main/rule-types/github/secret_scanning.yaml\ncurl -LO https://raw.githubusercontent.com/mindersec/minder-rules-and-profiles/main/rule-types/github/secret_push_protection.yaml\n"})}),"\n",(0,n.jsx)(r.p,{children:"Once you've downloaded the rule type configuration from GitHub, you can upload them to your Minder organization."}),"\n",(0,n.jsx)(r.pre,{children:(0,n.jsx)(r.code,{className:"language-bash",children:"minder ruletype create -f secret_scanning.yaml\nminder ruletype create -f secret_push_protection.yaml\n"})}),"\n",(0,n.jsx)(r.h3,{id:"creating-a-profile-1",children:"Creating a profile"}),"\n",(0,n.jsx)(r.p,{children:"Once you have added the rule type definitions to Minder, you can create a profile that uses them to scan your repositories."}),"\n",(0,n.jsx)(r.p,{children:"Like rules, profiles are defined in YAML. They specify the rules that apply to your organization, whether you want to be alerted with GitHub Security Advisories, and whether you want Minder to try to automatically remediate problems when they're found."}),"\n",(0,n.jsxs)(r.p,{children:["To create a profile that checks for secret scanning and secret push protection in your repositories, create a file called ",(0,n.jsx)(r.code,{children:"my_profile.yaml"}),":"]}),"\n",(0,n.jsx)(r.pre,{children:(0,n.jsx)(r.code,{className:"language-yaml",children:'---\nversion: v1\ntype: profile\nname: my_profile\ncontext:\n  provider: github\nalert: "on"\nremediate: "off"\nrepository:\n  - type: secret_scanning\n    def:\n      enabled: true\n  - type: secret_push_protection\n    def:\n      enabled: true\n'})}),"\n",(0,n.jsx)(r.p,{children:"Then upload the profile configuration to Minder:"}),"\n",(0,n.jsx)(r.pre,{children:(0,n.jsx)(r.code,{children:"minder profile create -f my_profile.yaml\n"})}),"\n",(0,n.jsx)(r.p,{children:"Check the status of the profile:"}),"\n",(0,n.jsx)(r.pre,{children:(0,n.jsx)(r.code,{children:"minder profile status list --name my_profile\n"})}),"\n",(0,n.jsxs)(r.p,{children:["If all registered repositories have secret scanning enabled, you will see the ",(0,n.jsx)(r.code,{children:"OVERALL STATUS"})," is ",(0,n.jsx)(r.code,{children:"Success"}),", otherwise the\noverall status is ",(0,n.jsx)(r.code,{children:"Failure"}),"."]}),"\n",(0,n.jsx)(r.pre,{children:(0,n.jsx)(r.code,{children:"+--------------------------------------+------------+----------------+----------------------+\n|                  ID                  |    NAME    | OVERALL STATUS |     LAST UPDATED     |\n+--------------------------------------+------------+----------------+----------------------+\n| 1abcae55-5eb8-4d9e-847c-18e605fbc1cc | my_profile |    Success     | 2023-11-06T17:42:04Z |\n+--------------------------------------+------------+----------------+----------------------+\n"})}),"\n",(0,n.jsxs)(r.p,{children:["If secret scanning is not enabled, you will see ",(0,n.jsx)(r.code,{children:"Failure"})," instead of ",(0,n.jsx)(r.code,{children:"Success"}),"."]}),"\n",(0,n.jsx)(r.p,{children:"See a detailed view of which repositories satisfy the secret scanning rule:"}),"\n",(0,n.jsx)(r.pre,{children:(0,n.jsx)(r.code,{children:"minder profile status list --name my_profile --detailed\n"})})]})}function p(e={}){const{wrapper:r}={...(0,i.R)(),...e.components};return r?(0,n.jsx)(r,{...e,children:(0,n.jsx)(d,{...e})}):d(e)}},28453:(e,r,t)=>{t.d(r,{R:()=>o,x:()=>a});var n=t(96540);const i={},s=n.createContext(i);function o(e){const r=n.useContext(s);return n.useMemo((function(){return"function"==typeof e?e(r):{...r,...e}}),[r,e])}function a(e){let r;return r=e.disableParentContext?"function"==typeof e.components?e.components(i):e.components||i:o(e.components),n.createElement(s.Provider,{value:r},e.children)}}}]);