"use strict";(self.webpackChunkminder_docs=self.webpackChunkminder_docs||[]).push([[4597],{37104:(e,i,t)=>{t.r(i),t.d(i,{assets:()=>c,contentTitle:()=>o,default:()=>u,frontMatter:()=>n,metadata:()=>a,toc:()=>l});var r=t(74848),s=t(28453);const n={title:"Viewing profile status",sidebar_position:70},o="Viewing the status of your profile",a={id:"getting_started/viewing_status",title:"Viewing profile status",description:"When you have created a profile and registered repositories, Minder will evaluate your security profile against those repositories. Minder can report on the status, and can optionally alert you using GitHub Security Advisories.",source:"@site/docs/getting_started/viewing_status.md",sourceDirName:"getting_started",slug:"/getting_started/viewing_status",permalink:"/getting_started/viewing_status",draft:!1,unlisted:!1,tags:[],version:"current",sidebarPosition:70,frontMatter:{title:"Viewing profile status",sidebar_position:70},sidebar:"minder",previous:{title:"Creating your first profile",permalink:"/getting_started/first_profile"},next:{title:"Automatic remediations",permalink:"/getting_started/remediations"}},c={},l=[{value:"Prerequisites",id:"prerequisites",level:2},{value:"Summary profile status",id:"summary-profile-status",level:2},{value:"Detailed profile status",id:"detailed-profile-status",level:2},{value:"Alerts with GitHub Security Advisories",id:"alerts-with-github-security-advisories",level:2}];function d(e){const i={a:"a",code:"code",em:"em",h1:"h1",h2:"h2",header:"header",p:"p",pre:"pre",...(0,s.R)(),...e.components};return(0,r.jsxs)(r.Fragment,{children:[(0,r.jsx)(i.header,{children:(0,r.jsx)(i.h1,{id:"viewing-the-status-of-your-profile",children:"Viewing the status of your profile"})}),"\n",(0,r.jsx)(i.p,{children:"When you have created a profile and registered repositories, Minder will evaluate your security profile against those repositories. Minder can report on the status, and can optionally alert you using GitHub Security Advisories."}),"\n",(0,r.jsx)(i.h2,{id:"prerequisites",children:"Prerequisites"}),"\n",(0,r.jsxs)(i.p,{children:["Before you can view the status of your profile, you must ",(0,r.jsx)(i.a,{href:"first_profile",children:"create a profile"}),"."]}),"\n",(0,r.jsx)(i.h2,{id:"summary-profile-status",children:"Summary profile status"}),"\n",(0,r.jsxs)(i.p,{children:["To view the summary status of your profile, use the ",(0,r.jsx)(i.code,{children:"minder profile status list"})," command. If you have a profile named ",(0,r.jsx)(i.code,{children:"my_profile"}),", run:"]}),"\n",(0,r.jsx)(i.pre,{children:(0,r.jsx)(i.code,{className:"language-bash",children:"minder profile status list --name my_profile\n"})}),"\n",(0,r.jsxs)(i.p,{children:["If all the registered repositories are in compliance with your profile, you will see the ",(0,r.jsx)(i.code,{children:"OVERALL STATUS"})," column set to ",(0,r.jsx)(i.code,{children:"Success"}),". If one or more repositories are not in compliance, the column will be set to ",(0,r.jsx)(i.code,{children:"Failure"}),"."]}),"\n",(0,r.jsxs)(i.p,{children:["For example, the profile named ",(0,r.jsx)(i.code,{children:"my_profile"})," expects repositories to have secret scanning enabled. If any repository did not have secret scanning enabled, then the output will look like:"]}),"\n",(0,r.jsx)(i.pre,{children:(0,r.jsx)(i.code,{className:"language-yaml",children:"+--------------------------------------+------------+----------------+----------------------+\n|                  ID                  |    NAME    | OVERALL STATUS |     LAST UPDATED     |\n+--------------------------------------+------------+----------------+----------------------+\n| 1abcae55-5eb8-4d9e-847c-18e605fbc1cc | my_profile |    Failed      | 2023-11-06T17:42:04Z |\n+--------------------------------------+------------+----------------+----------------------+\n"})}),"\n",(0,r.jsx)(i.p,{children:"Use detailed status reporting to understand which repositories are not in compliance."}),"\n",(0,r.jsx)(i.h2,{id:"detailed-profile-status",children:"Detailed profile status"}),"\n",(0,r.jsx)(i.p,{children:"Detailed status will show each repository that is registered, along with the current evaluation status for each rule."}),"\n",(0,r.jsx)(i.p,{children:"See a detailed view of which repositories satisfy the secret scanning rule:"}),"\n",(0,r.jsx)(i.pre,{children:(0,r.jsx)(i.code,{className:"language-bash",children:"minder profile status list --name github-profile --detailed\n"})}),"\n",(0,r.jsxs)(i.p,{children:["An example output for a profile that checks secret scanning and secret push protection, for an organization that has a single repository registered. In this example, the repository ",(0,r.jsx)(i.code,{children:"example/demo_repo"})," has secret scanning enabled, which is indicated by the ",(0,r.jsx)(i.code,{children:"STATUS"})," column set to ",(0,r.jsx)(i.code,{children:"Success"}),". However, that repository does not have secret push protection enabled, which is indicated by the ",(0,r.jsx)(i.code,{children:"STATUS"})," column set to ",(0,r.jsx)(i.code,{children:"Failure"}),"."]}),"\n",(0,r.jsx)(i.pre,{children:(0,r.jsx)(i.code,{className:"language-yaml",children:"+--------------------------------------+------------------------+------------+---------+-------------+--------------------------------------+\n|               RULE ID                |       RULE NAME        |   ENTITY   | STATUS  | REMEDIATION |             ENTITY INFO              |\n+--------------------------------------+------------------------+------------+---------+-------------+--------------------------------------+\n| 8a2af1c3-72f6-42ac-a888-45eac5b0f72e | secret_scanning        | repository | Success | Skipped     | provider: github-app-example         |\n|                                      |                        |            |         |             | repo_name: demo_repo repo_owner:     |\n|                                      |                        |            |         |             | example  repository_id:              |\n|                                      |                        |            |         |             | 04055a1a-766e-4f49-a1ba-16ab1e749fef |\n|                                      |                        |            |         |             |                                      |\n+--------------------------------------+------------------------+------------+---------+-------------+--------------------------------------+\n| 08e94b93-e3d6-4df5-a480-ecf108ba481e | secret_push_protection | repository | Failure | Skipped     | provider: github-app-example         |\n|                                      |                        |            |         |             | repo_name: demo_repo repo_owner:     |\n|                                      |                        |            |         |             | example  repository_id:              |\n|                                      |                        |            |         |             | 04055a1a-766e-4f49-a1ba-16ab1e749fef |\n|                                      |                        |            |         |             |                                      |\n+--------------------------------------+------------------------+------------+---------+-------------+--------------------------------------+\n"})}),"\n",(0,r.jsx)(i.h2,{id:"alerts-with-github-security-advisories",children:"Alerts with GitHub Security Advisories"}),"\n",(0,r.jsxs)(i.p,{children:["You can optionally get alerted with ",(0,r.jsx)(i.a,{href:"https://docs.github.com/en/code-security/security-advisories",children:"GitHub Security Advisories"})," when repositories are not in compliance with your security profiles. If you have configured your profile with ",(0,r.jsx)(i.code,{children:"alerts: on"}),", then Minder will generate GitHub Security Advisories."]}),"\n",(0,r.jsxs)(i.p,{children:["For example, if you've ",(0,r.jsxs)(i.a,{href:"first_profile",children:["created a profile with ",(0,r.jsx)(i.code,{children:"alerts: on"})]})," that looks for secret scanning to be enabled in your repository, then ",(0,r.jsx)(i.em,{children:"disabling"})," secret scanning in that repository should produce a GitHub Security Advisory."]}),"\n",(0,r.jsxs)(i.p,{children:["In this example, if you ",(0,r.jsx)(i.a,{href:"https://docs.github.com/en/code-security/secret-scanning/configuring-secret-scanning-for-your-repositories",children:"disable secret scanning"})," in one of your registered repositories, Minder will create a GitHub Security Advisory in that repository. To view that, navigate to the repository on GitHub, click on the Security tab and view the Security Advisories. There will be a new advisories named ",(0,r.jsx)(i.code,{children:"minder: profile my_profile failed with rule secret_scanning"}),"."]}),"\n",(0,r.jsxs)(i.p,{children:["To resolve this issue, you can ",(0,r.jsx)(i.a,{href:"https://docs.github.com/en/code-security/secret-scanning/configuring-secret-scanning-for-your-repositories",children:"enable secret scanning"})," in that repository. When you do this, the advisory will be deleted. If you go back to the Security Advisories page on that repository, you will see that the advisory that was created by Minder has been closed."]})]})}function u(e={}){const{wrapper:i}={...(0,s.R)(),...e.components};return i?(0,r.jsx)(i,{...e,children:(0,r.jsx)(d,{...e})}):d(e)}},28453:(e,i,t)=>{t.d(i,{R:()=>o,x:()=>a});var r=t(96540);const s={},n=r.createContext(s);function o(e){const i=r.useContext(n);return r.useMemo((function(){return"function"==typeof e?e(i):{...i,...e}}),[i,e])}function a(e){let i;return i=e.disableParentContext?"function"==typeof e.components?e.components(s):e.components||s:o(e.components),r.createElement(n.Provider,{value:i},e.children)}}}]);