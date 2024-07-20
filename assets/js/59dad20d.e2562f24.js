"use strict";(self.webpackChunkminder_docs=self.webpackChunkminder_docs||[]).push([[486],{60488:(e,n,i)=>{i.r(n),i.d(n,{assets:()=>l,contentTitle:()=>a,default:()=>p,frontMatter:()=>s,metadata:()=>o,toc:()=>d});var t=i(74848),r=i(28453);const s={title:"Profiles and Rules",sidebar_position:10},a="Profiles in Minder",o={id:"understand/profiles",title:"Profiles and Rules",description:"A profile defines your security policies that you want to apply to your software supply chain. Profiles contain rules that query data in a provider, and specifies whether Minder will issue alerts or perform automatic remediations when an entity is not in compliance with the policy.",source:"@site/docs/understand/profiles.md",sourceDirName:"understand",slug:"/understand/profiles",permalink:"/understand/profiles",draft:!1,unlisted:!1,tags:[],version:"current",sidebarPosition:10,frontMatter:{title:"Profiles and Rules",sidebar_position:10},sidebar:"minder",previous:{title:"Automatic remediations",permalink:"/getting_started/remediations"},next:{title:"Projects",permalink:"/understand/projects"}},l={},d=[{value:"Rules",id:"rules",level:2},{value:"Actions",id:"actions",level:2},{value:"Example profile",id:"example-profile",level:2}];function c(e){const n={a:"a",code:"code",em:"em",h1:"h1",h2:"h2",li:"li",p:"p",pre:"pre",strong:"strong",ul:"ul",...(0,r.R)(),...e.components};return(0,t.jsxs)(t.Fragment,{children:[(0,t.jsx)(n.h1,{id:"profiles-in-minder",children:"Profiles in Minder"}),"\n",(0,t.jsxs)(n.p,{children:["A ",(0,t.jsx)(n.em,{children:"profile"})," defines your security policies that you want to apply to your software supply chain. Profiles contain rules that query data in a ",(0,t.jsx)(n.a,{href:"/understand/providers",children:"provider"}),", and specifies whether Minder will issue ",(0,t.jsx)(n.a,{href:"/understand/alerts",children:"alerts"})," or perform automatic ",(0,t.jsx)(n.a,{href:"/understand/remediations",children:"remediations"})," when an entity is not in compliance with the policy."]}),"\n",(0,t.jsx)(n.p,{children:"Profiles in Minder allow you to group and manage\nrules for various entity types, such as repositories, pull requests, and artifacts, across your registered GitHub\nrepositories."}),"\n",(0,t.jsx)(n.p,{children:"The anatomy of a profile is the profile itself, which outlines the rules to be\nchecked, the rule types, and the evaluation engine. Profiles are defined in a YAML file and can be configured to meet your compliance requirements.\nAs of time of writing, Minder supports the following evaluation engines:"}),"\n",(0,t.jsxs)(n.ul,{children:["\n",(0,t.jsxs)(n.li,{children:[(0,t.jsx)(n.strong,{children:(0,t.jsx)(n.a,{href:"https://www.openpolicyagent.org/docs/latest/policy-language/",children:"Rego"})})," Open Policy Agents's native query language, Rego."]}),"\n",(0,t.jsxs)(n.li,{children:[(0,t.jsx)(n.strong,{children:(0,t.jsx)(n.a,{href:"https://jqlang.github.io/jq/",children:"JQ"})})," - a lightweight and flexible command-line JSON processor."]}),"\n"]}),"\n",(0,t.jsx)(n.p,{children:"Each engine is designed to be extensible, allowing you to integrate your own\nlogic and processes."}),"\n",(0,t.jsxs)(n.p,{children:["Stacklok has published ",(0,t.jsx)(n.a,{href:"https://github.com/stacklok/minder-rules-and-profiles/tree/main/profiles/github",children:"a set of example profiles on GitHub"}),"; see ",(0,t.jsx)(n.a,{href:"/how-to/manage_profiles",children:"Managing Profiles"})," for more details on how to set up profiles and rule types."]}),"\n",(0,t.jsx)(n.h2,{id:"rules",children:"Rules"}),"\n",(0,t.jsx)(n.p,{children:"Each rule type within a profile is evaluated against your repositories that are registered with Minder."}),"\n",(0,t.jsxs)(n.p,{children:["The available entity rule type groups are ",(0,t.jsx)(n.code,{children:"repository"}),", ",(0,t.jsx)(n.code,{children:"pull_request"}),", and ",(0,t.jsx)(n.code,{children:"artifact"}),"."]}),"\n",(0,t.jsx)(n.p,{children:"Each rule type group has a set of rules that can be configured individually."}),"\n",(0,t.jsxs)(n.p,{children:["The available properties for each rule type can be found in their YAML file under the ",(0,t.jsx)(n.code,{children:"def.rule_schema.properties"})," section."]}),"\n",(0,t.jsxs)(n.p,{children:["For example, the ",(0,t.jsx)(n.code,{children:"secret_scanning"})," rule type has the following ",(0,t.jsx)(n.code,{children:"enabled"})," property:"]}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-yaml",children:"---\nversion: v1\ntype: rule-type\nname: secret_scanning\n---\ndef:\n  # Defines the schema for writing a rule with this rule being checked\n  rule_schema:\n    properties:\n      enabled:\n        type: boolean\n        default: true\n"})}),"\n",(0,t.jsxs)(n.p,{children:["You can find a list of available rules in the ",(0,t.jsx)(n.a,{href:"https://github.com/stacklok/minder-rules-and-profiles/tree/main/rule-types/github",children:"https://github.com/stacklok/minder-rules-and-profiles"})," repository."]}),"\n",(0,t.jsx)(n.h2,{id:"actions",children:"Actions"}),"\n",(0,t.jsx)(n.p,{children:"Minder supports the ability to perform actions based on the evaluation of a rule such as creating alerts\nand automatically remediating non-compliant rules."}),"\n",(0,t.jsx)(n.p,{children:"When a rule fails, Minder can open an alert to bring your attention to the non-compliance issue. Conversely, when the\nrule evaluation passes, Minder will automatically close any previously opened alerts related to that rule."}),"\n",(0,t.jsx)(n.p,{children:"Minder also supports the ability to automatically remediate failed rules based on their type, i.e., by processing a\nREST call to enable/disable a non-compliant repository setting or creating a pull request with a proposed fix. Note\nthat not all rule types support automatic remediation yet."}),"\n",(0,t.jsxs)(n.p,{children:["Both alerts and remediations are configured in the profile YAML file under ",(0,t.jsx)(n.code,{children:"alerts"})," (Default: ",(0,t.jsx)(n.code,{children:"on"}),")\nand ",(0,t.jsx)(n.code,{children:"remediate"})," (Default: ",(0,t.jsx)(n.code,{children:"off"}),")."]}),"\n",(0,t.jsx)(n.h2,{id:"example-profile",children:"Example profile"}),"\n",(0,t.jsxs)(n.p,{children:["Here's a profile which has a single rule for each entity group and its ",(0,t.jsx)(n.code,{children:"alert"})," and ",(0,t.jsx)(n.code,{children:"remediate"})," features are both\nturned ",(0,t.jsx)(n.code,{children:"on"}),":"]}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-yaml",children:'---\nversion: v1\ntype: profile\nname: github-profile\ncontext:\n  provider: github\nalert: "on"\nremediate: "on"\nrepository:\n  - type: dependabot_configured\n    def:\n      package_ecosystem: gomod\n      schedule_interval: daily\n      apply_if_file: go.mod\nartifact:\n  - type: artifact_signature\n    params:\n      tags: [latest]\n      name: your-artifact-name\n    def:\n      is_signed: true\n      is_verified: true\npull_request:\n  - type: pr_vulnerability_check\n    def:\n      action: review\n      ecosystem_config:\n        - name: npm\n          vulnerability_database_type: osv\n          vulnerability_database_endpoint: https://api.osv.dev/v1/query\n          package_repository:\n            url: https://registry.npmjs.org\n        - name: go\n          vulnerability_database_type: osv\n          vulnerability_database_endpoint: https://api.osv.dev/v1/query\n          package_repository:\n            url: https://proxy.golang.org\n          sum_repository:\n            url: https://sum.golang.org\n        - name: pypi\n          vulnerability_database_type: osv\n          vulnerability_database_endpoint: https://api.osv.dev/v1/query\n          package_repository:\n            url: https://pypi.org/pypi\n'})})]})}function p(e={}){const{wrapper:n}={...(0,r.R)(),...e.components};return n?(0,t.jsx)(n,{...e,children:(0,t.jsx)(c,{...e})}):c(e)}},28453:(e,n,i)=>{i.d(n,{R:()=>a,x:()=>o});var t=i(96540);const r={},s=t.createContext(r);function a(e){const n=t.useContext(s);return t.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function o(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(r):e.components||r:a(e.components),t.createElement(s.Provider,{value:n},e.children)}}}]);