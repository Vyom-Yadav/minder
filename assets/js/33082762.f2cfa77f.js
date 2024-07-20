"use strict";(self.webpackChunkminder_docs=self.webpackChunkminder_docs||[]).push([[2828],{30323:(e,t,i)=>{i.r(t),i.d(t,{assets:()=>a,contentTitle:()=>s,default:()=>h,frontMatter:()=>l,metadata:()=>o,toc:()=>d});var r=i(74848),n=i(28453);const l={title:"Setting up a profile for GitHub Security Advisories",sidebar_position:70},s="Setting up a profile for GitHub Security Advisories",o={id:"how-to/setup-alerts",title:"Setting up a profile for GitHub Security Advisories",description:"Prerequisites",source:"@site/docs/how-to/setup-alerts.md",sourceDirName:"how-to",slug:"/how-to/setup-alerts",permalink:"/how-to/setup-alerts",draft:!1,unlisted:!1,tags:[],version:"current",sidebarPosition:70,frontMatter:{title:"Setting up a profile for GitHub Security Advisories",sidebar_position:70},sidebar:"minder",previous:{title:"Automatic remediation via Pull Request",permalink:"/how-to/remediate-pullrequest"},next:{title:"Writing custom rule types",permalink:"/how-to/custom-rules"}},a={},d=[{value:"Prerequisites",id:"prerequisites",level:2},{value:"Create a rule type that you want to be alerted on",id:"create-a-rule-type-that-you-want-to-be-alerted-on",level:2},{value:"Create a profile",id:"create-a-profile",level:2},{value:"Limitations",id:"limitations",level:2}];function c(e){const t={a:"a",code:"code",h1:"h1",h2:"h2",li:"li",p:"p",pre:"pre",ul:"ul",...(0,n.R)(),...e.components};return(0,r.jsxs)(r.Fragment,{children:[(0,r.jsx)(t.h1,{id:"setting-up-a-profile-for-github-security-advisories",children:"Setting up a profile for GitHub Security Advisories"}),"\n",(0,r.jsx)(t.h2,{id:"prerequisites",children:"Prerequisites"}),"\n",(0,r.jsxs)(t.ul,{children:["\n",(0,r.jsxs)(t.li,{children:["The ",(0,r.jsx)(t.code,{children:"minder"})," CLI application"]}),"\n",(0,r.jsx)(t.li,{children:"A Minder account"}),"\n",(0,r.jsx)(t.li,{children:"An enrolled Provider (e.g., GitHub) and registered repositories"}),"\n"]}),"\n",(0,r.jsx)(t.h2,{id:"create-a-rule-type-that-you-want-to-be-alerted-on",children:"Create a rule type that you want to be alerted on"}),"\n",(0,r.jsxs)(t.p,{children:["The ",(0,r.jsx)(t.code,{children:"alert"})," feature is available for all rule types that have the ",(0,r.jsx)(t.code,{children:"alert"})," section defined in their ",(0,r.jsx)(t.code,{children:"<alert-type>.yaml"}),"\nfile. Alerts are a core feature of Minder providing you with notifications about the status of your registered\nrepositories using GitHub Security Advisories. These Security Advisories automatically open and close based on the evaluation of the rules defined in your profiles."]}),"\n",(0,r.jsx)(t.p,{children:"When a rule fails, Minder opens an alert to bring your attention to the non-compliance issue. Conversely, when the\nrule evaluation passes, Minder will automatically close any previously opened alerts related to that rule."}),"\n",(0,r.jsxs)(t.p,{children:["In this example, we will use a rule type that checks if a repository has a LICENSE file present. If there's no file\npresent, Minder will create an alert notifying the owner of the repository. The rule type is called ",(0,r.jsx)(t.code,{children:"license.yaml"})," and\nis one of the reference rule types provided by the Minder team. Details, such as the severity of the alert are defined\nin the ",(0,r.jsx)(t.code,{children:"alert"})," section of the rule type definition."]}),"\n",(0,r.jsxs)(t.p,{children:["Fetch all the reference rules by cloning the ",(0,r.jsx)(t.a,{href:"https://github.com/stacklok/minder-rules-and-profiles",children:"minder-rules-and-profiles repository"}),"."]}),"\n",(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-bash",children:"git clone https://github.com/stacklok/minder-rules-and-profiles.git\n"})}),"\n",(0,r.jsx)(t.p,{children:"In that directory, you can find all the reference rules and profiles."}),"\n",(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-bash",children:"cd minder-rules-and-profiles\n"})}),"\n",(0,r.jsxs)(t.p,{children:["Create the ",(0,r.jsx)(t.code,{children:"license"})," rule type in Minder:"]}),"\n",(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-bash",children:"minder ruletype create -f rule-types/github/license.yaml\n"})}),"\n",(0,r.jsx)(t.h2,{id:"create-a-profile",children:"Create a profile"}),"\n",(0,r.jsx)(t.p,{children:"Next, create a profile that applies the rule to all registered repositories."}),"\n",(0,r.jsxs)(t.p,{children:["Create a new file called ",(0,r.jsx)(t.code,{children:"profile.yaml"})," using the following profile definition and enable alerting by setting ",(0,r.jsx)(t.code,{children:"alert"}),"\nto ",(0,r.jsx)(t.code,{children:"on"})," (default). The other available values are ",(0,r.jsx)(t.code,{children:"off"})," and ",(0,r.jsx)(t.code,{children:"dry_run"}),"."]}),"\n",(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-yaml",children:'---\nversion: v1\ntype: profile\nname: license-profile\ncontext:\n  provider: github\nalert: "on"\nrepository:\n  - type: license\n    def:\n      license_filename: LICENSE\n      license_type: ""\n'})}),"\n",(0,r.jsx)(t.p,{children:"Create the profile in Minder:"}),"\n",(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-bash",children:"minder profile create -f profile.yaml\n"})}),"\n",(0,r.jsxs)(t.p,{children:["Once the profile is created, Minder will monitor all of your registered repositories for the presence of the ",(0,r.jsx)(t.code,{children:"LICENSE"}),"\nfile."]}),"\n",(0,r.jsxs)(t.p,{children:["If a repository does not have a ",(0,r.jsx)(t.code,{children:"LICENSE"})," file available, Minder will create an alert of type Security Advisory providing\nadditional details such as the profile and rule that triggered the alert and guidelines on how to resolve the issue."]}),"\n",(0,r.jsxs)(t.p,{children:["Once a ",(0,r.jsx)(t.code,{children:"LICENSE"})," file is added to the repository, Minder will automatically close the alert."]}),"\n",(0,r.jsxs)(t.p,{children:["Alerts are complementary to the remediation feature. If you have both ",(0,r.jsx)(t.code,{children:"alert"})," and ",(0,r.jsx)(t.code,{children:"remediation"})," enabled for a profile,\nMinder will attempt to remediate it first. If the remediation fails, Minder will create an alert. If the remediation\nsucceeds, Minder will close any previously opened alerts related to that rule."]}),"\n",(0,r.jsx)(t.h2,{id:"limitations",children:"Limitations"}),"\n",(0,r.jsxs)(t.ul,{children:["\n",(0,r.jsx)(t.li,{children:"Currently, the only supported alert type is GitHub Security Advisory. More alert types will be added in the future."}),"\n",(0,r.jsxs)(t.li,{children:["Alerts are only available for rules that have the ",(0,r.jsx)(t.code,{children:"alert"})," section defined in their ",(0,r.jsx)(t.code,{children:"<alert-type>.yaml"})," file."]}),"\n"]})]})}function h(e={}){const{wrapper:t}={...(0,n.R)(),...e.components};return t?(0,r.jsx)(t,{...e,children:(0,r.jsx)(c,{...e})}):c(e)}},28453:(e,t,i)=>{i.d(t,{R:()=>s,x:()=>o});var r=i(96540);const n={},l=r.createContext(n);function s(e){const t=r.useContext(l);return r.useMemo((function(){return"function"==typeof e?e(t):{...t,...e}}),[t,e])}function o(e){let t;return t=e.disableParentContext?"function"==typeof e.components?e.components(n):e.components||n:s(e.components),r.createElement(l.Provider,{value:t},e.children)}}}]);