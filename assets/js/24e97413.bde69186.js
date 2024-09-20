"use strict";(self.webpackChunkminder_docs=self.webpackChunkminder_docs||[]).push([[5882],{43778:(e,n,r)=>{r.r(n),r.d(n,{assets:()=>c,contentTitle:()=>o,default:()=>h,frontMatter:()=>i,metadata:()=>l,toc:()=>d});var s=r(74848),t=r(28453);const i={title:"Trusty Score",sidebar_position:20},o="Trusty Score Threshold Rule",l={id:"ref/rules/pr_trusty_check",title:"Trusty Score",description:"The following rule type is available for Trusty score threshold.",source:"@site/docs/ref/rules/pr_trusty_check.md",sourceDirName:"ref/rules",slug:"/ref/rules/pr_trusty_check",permalink:"/ref/rules/pr_trusty_check",draft:!1,unlisted:!1,tags:[],version:"current",sidebarPosition:20,frontMatter:{title:"Trusty Score",sidebar_position:20},sidebar:"minder",previous:{title:"Branch protections",permalink:"/ref/rules/branch_protection"},next:{title:"Dependabot",permalink:"/ref/rules/dependabot_configured"}},c={},d=[{value:"<code>pr_trusty_check</code> - Verifies that pull requests do not add any dependencies with Trusty scores below a certain threshold",id:"pr_trusty_check---verifies-that-pull-requests-do-not-add-any-dependencies-with-trusty-scores-below-a-certain-threshold",level:2},{value:"Entity",id:"entity",level:2},{value:"Type",id:"type",level:2},{value:"Rule Parameters",id:"rule-parameters",level:2},{value:"Rule Definition Options",id:"rule-definition-options",level:2}];function a(e){const n={a:"a",code:"code",h1:"h1",h2:"h2",header:"header",li:"li",p:"p",ul:"ul",...(0,t.R)(),...e.components};return(0,s.jsxs)(s.Fragment,{children:[(0,s.jsx)(n.header,{children:(0,s.jsx)(n.h1,{id:"trusty-score-threshold-rule",children:"Trusty Score Threshold Rule"})}),"\n",(0,s.jsxs)(n.p,{children:["The following rule type is available for ",(0,s.jsx)(n.a,{href:"https://www.trustypkg.dev/",children:"Trusty"})," score threshold."]}),"\n",(0,s.jsxs)(n.h2,{id:"pr_trusty_check---verifies-that-pull-requests-do-not-add-any-dependencies-with-trusty-scores-below-a-certain-threshold",children:[(0,s.jsx)(n.code,{children:"pr_trusty_check"})," - Verifies that pull requests do not add any dependencies with Trusty scores below a certain threshold"]}),"\n",(0,s.jsxs)(n.p,{children:["This rule allows you to monitor new pull requests for newly added dependencies with low\n",(0,s.jsx)(n.a,{href:"https://www.trustypkg.dev/",children:"Trusty"})," scores.\nFor every pull request submitted to a repository, this rule will check if the pull request adds a new dependency with\na Trusty score below a threshold that you define. If a dependency with a low score is added, the PR will be commented on."]}),"\n",(0,s.jsx)(n.h2,{id:"entity",children:"Entity"}),"\n",(0,s.jsxs)(n.ul,{children:["\n",(0,s.jsx)(n.li,{children:(0,s.jsx)(n.code,{children:"pull_request"})}),"\n"]}),"\n",(0,s.jsx)(n.h2,{id:"type",children:"Type"}),"\n",(0,s.jsxs)(n.ul,{children:["\n",(0,s.jsx)(n.li,{children:(0,s.jsx)(n.code,{children:"pr_trusty_check"})}),"\n"]}),"\n",(0,s.jsx)(n.h2,{id:"rule-parameters",children:"Rule Parameters"}),"\n",(0,s.jsxs)(n.ul,{children:["\n",(0,s.jsx)(n.li,{children:"None"}),"\n"]}),"\n",(0,s.jsx)(n.h2,{id:"rule-definition-options",children:"Rule Definition Options"}),"\n",(0,s.jsxs)(n.p,{children:["The ",(0,s.jsx)(n.code,{children:"pr_trusty_check"})," rule has the following options:"]}),"\n",(0,s.jsxs)(n.ul,{children:["\n",(0,s.jsxs)(n.li,{children:[(0,s.jsx)(n.code,{children:"action"})," (string): The action to take if a package with a low score is found. Valid values are:","\n",(0,s.jsxs)(n.ul,{children:["\n",(0,s.jsxs)(n.li,{children:[(0,s.jsx)(n.code,{children:"summary"}),": The evaluator engine will add a single summary comment with a table listing the packages with low scores found"]}),"\n",(0,s.jsxs)(n.li,{children:[(0,s.jsx)(n.code,{children:"profile_only"}),": The evaluator engine will merely pass on an error, marking the profile as failed if a packages with low scores is found"]}),"\n",(0,s.jsxs)(n.li,{children:[(0,s.jsx)(n.code,{children:"review"}),": The trusty evaluator will add a review asking for changes when problematic dependencies are found. Use the review action to block any pull requests introducing dependencies that break the policy established defined by the rule."]}),"\n"]}),"\n"]}),"\n",(0,s.jsxs)(n.li,{children:[(0,s.jsx)(n.code,{children:"ecosystem_config"}),": An array of ecosystem configurations to check. Each ecosystem configuration has the following options:","\n",(0,s.jsxs)(n.ul,{children:["\n",(0,s.jsxs)(n.li,{children:[(0,s.jsx)(n.code,{children:"name"})," (string): The name of the ecosystem to check. Currently ",(0,s.jsx)(n.code,{children:"npm"})," and ",(0,s.jsx)(n.code,{children:"pypi"})," are supported."]}),"\n",(0,s.jsxs)(n.li,{children:[(0,s.jsx)(n.code,{children:"score"})," (number): The minimum Trusty score for a dependency to be considered safe."]}),"\n",(0,s.jsxs)(n.li,{children:[(0,s.jsx)(n.code,{children:"provenance"})," (number): Minimum provenance score to consider a package's proof of origin satisfactory."]}),"\n",(0,s.jsxs)(n.li,{children:[(0,s.jsx)(n.code,{children:"activity"})," (number): Minimum activity score to consider a package as active."]}),"\n",(0,s.jsxs)(n.li,{children:[(0,s.jsx)(n.code,{children:"allow_malicious"})," (boolean): Don't raise an error when a PR introduces dependencies known to be malicious (not recommended)"]}),"\n",(0,s.jsxs)(n.li,{children:[(0,s.jsx)(n.code,{children:"allow_deprecated"})," (boolean): Don't block when a pull request introduces dependencies marked as deprecated upstream."]}),"\n"]}),"\n"]}),"\n"]})]})}function h(e={}){const{wrapper:n}={...(0,t.R)(),...e.components};return n?(0,s.jsx)(n,{...e,children:(0,s.jsx)(a,{...e})}):a(e)}},28453:(e,n,r)=>{r.d(n,{R:()=>o,x:()=>l});var s=r(96540);const t={},i=s.createContext(t);function o(e){const n=s.useContext(i);return s.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function l(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(t):e.components||t:o(e.components),s.createElement(i.Provider,{value:n},e.children)}}}]);