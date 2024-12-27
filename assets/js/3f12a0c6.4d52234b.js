"use strict";(self.webpackChunkminder_docs=self.webpackChunkminder_docs||[]).push([[6777],{36938:(e,n,i)=>{i.r(n),i.d(n,{assets:()=>d,contentTitle:()=>c,default:()=>h,frontMatter:()=>s,metadata:()=>t,toc:()=>a});const t=JSON.parse('{"id":"user_management/adding_users","title":"Adding Users to your Project","description":"Prerequisites","source":"@site/docs/user_management/adding_users.md","sourceDirName":"user_management","slug":"/user_management/adding_users","permalink":"/user_management/adding_users","draft":false,"unlisted":false,"tags":[],"version":"current","sidebarPosition":100,"frontMatter":{"title":"Adding Users to your Project","sidebar_position":100},"sidebar":"minder","previous":{"title":"Feature flags","permalink":"/developer_guide/feature_flags"},"next":{"title":"User roles in Minder","permalink":"/user_management/user_roles"}}');var r=i(74848),o=i(28453);const s={title:"Adding Users to your Project",sidebar_position:100},c=void 0,d={},a=[{value:"Prerequisites",id:"prerequisites",level:2},{value:"Overview",id:"overview",level:2},{value:"Identify their role",id:"identify-their-role",level:2},{value:"Invite the user to your project",id:"invite-the-user-to-your-project",level:2},{value:"Have the User Exercise the Invitation Code",id:"have-the-user-exercise-the-invitation-code",level:2},{value:"Viewing outstanding invitations",id:"viewing-outstanding-invitations",level:2},{value:"Working with Multiple Projects",id:"working-with-multiple-projects",level:2}];function l(e){const n={a:"a",code:"code",em:"em",h2:"h2",li:"li",ol:"ol",p:"p",pre:"pre",strong:"strong",ul:"ul",...(0,o.R)(),...e.components};return(0,r.jsxs)(r.Fragment,{children:[(0,r.jsx)(n.h2,{id:"prerequisites",children:"Prerequisites"}),"\n",(0,r.jsxs)(n.ul,{children:["\n",(0,r.jsxs)(n.li,{children:["The ",(0,r.jsx)(n.code,{children:"minder"})," CLI application"]}),"\n",(0,r.jsxs)(n.li,{children:["A Minder account with ",(0,r.jsxs)(n.a,{href:"/user_management/user_roles",children:[(0,r.jsx)(n.code,{children:"admin"})," permission"]})]}),"\n"]}),"\n",(0,r.jsx)(n.h2,{id:"overview",children:"Overview"}),"\n",(0,r.jsx)(n.p,{children:"To invite a new user to your project, you need to:"}),"\n",(0,r.jsxs)(n.ol,{children:["\n",(0,r.jsx)(n.li,{children:"Identify the new user's role"}),"\n",(0,r.jsx)(n.li,{children:"Invite the user to your project"}),"\n",(0,r.jsx)(n.li,{children:"Have the user exercise the invitation code"}),"\n"]}),"\n",(0,r.jsx)(n.h2,{id:"identify-their-role",children:"Identify their role"}),"\n",(0,r.jsx)(n.p,{children:"When adding a user into your project, it's crucial to assign them the\nappropriate role based on their responsibilities and required access levels."}),"\n",(0,r.jsxs)(n.p,{children:["Roles are ",(0,r.jsx)(n.a,{href:"/user_management/user_roles",children:"documented here"}),". To view the available roles in your\nproject, and their descriptions, run:"]}),"\n",(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{children:"minder project role list\n"})}),"\n",(0,r.jsx)(n.h2,{id:"invite-the-user-to-your-project",children:"Invite the user to your project"}),"\n",(0,r.jsxs)(n.p,{children:["Minder uses an invitation code system to allow users to accept (or decline) the\ninvitation to join a project. The invitation code may be delivered over email,\nor may be copied into a channel like Slack or a ticket system. ",(0,r.jsx)(n.strong,{children:"Any user with\naccess to an unused invitation code may accept the invitation, so treat the code\nlike a password."})]}),"\n",(0,r.jsx)(n.p,{children:"To add a user to your project, follow these steps:"}),"\n",(0,r.jsxs)(n.ol,{children:["\n",(0,r.jsxs)(n.li,{children:["\n",(0,r.jsx)(n.p,{children:"Determine the User's Role: Decide the appropriate role for the user based on\ntheir responsibilities."}),"\n"]}),"\n",(0,r.jsxs)(n.li,{children:["\n",(0,r.jsx)(n.p,{children:"Execute the Command:"}),"\n",(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-bash",children:"minder project role grant --email email@example.com --role desired-role\n"})}),"\n",(0,r.jsxs)(n.ul,{children:["\n",(0,r.jsxs)(n.li,{children:["Replace ",(0,r.jsx)(n.code,{children:"email@example.com"})," with the email address of the user you want to\ninvite."]}),"\n",(0,r.jsxs)(n.li,{children:["Replace ",(0,r.jsx)(n.code,{children:"desired-role"})," with the chosen role for the user (e.g., ",(0,r.jsx)(n.code,{children:"viewer"}),",\n",(0,r.jsx)(n.code,{children:"editor"}),")."]}),"\n"]}),"\n"]}),"\n",(0,r.jsxs)(n.li,{children:["\n",(0,r.jsxs)(n.p,{children:["You will receive a response message including an invitation code which can be\nused with\n",(0,r.jsx)(n.a,{href:"/ref/cli/minder_auth_invite_accept",children:(0,r.jsx)(n.code,{children:"minder auth invite accept"})}),"."]}),"\n"]}),"\n"]}),"\n",(0,r.jsx)(n.h2,{id:"have-the-user-exercise-the-invitation-code",children:"Have the User Exercise the Invitation Code"}),"\n",(0,r.jsxs)(n.p,{children:["Relay the invitation code to the user who you are inviting to join your project.\nThey will need to install the ",(0,r.jsx)(n.code,{children:"minder"})," CLI and run\n",(0,r.jsx)(n.code,{children:"minder auth invite accept <invitation code>"})," to accept the invitation.\nInvitations will expire when used, or after 7 days, though users who have not\naccepted an invitation can be invited again."]}),"\n",(0,r.jsx)(n.h2,{id:"viewing-outstanding-invitations",children:"Viewing outstanding invitations"}),"\n",(0,r.jsx)(n.p,{children:"You can then view all the project collaborators and outstanding user invitations\nby executing:"}),"\n",(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-bash",children:"minder project role grant list\n"})}),"\n",(0,r.jsx)(n.h2,{id:"working-with-multiple-projects",children:"Working with Multiple Projects"}),"\n",(0,r.jsxs)(n.p,{children:["When you have access to more than one project, you will need to qualify many\n",(0,r.jsx)(n.code,{children:"minder"})," commands with which project you want them to apply to. You can either\nuse the ",(0,r.jsx)(n.code,{children:"--project"})," flag, set a default in your\n",(0,r.jsx)(n.a,{href:"/ref/cli_configuration",children:"minder configuration"}),", or set the ",(0,r.jsx)(n.code,{children:"MINDER_PROJECT"}),"\nenvironment variable. To see all the projects that are available to you, use the\n",(0,r.jsx)(n.a,{href:"/ref/cli/minder_project_list",children:(0,r.jsx)(n.code,{children:"minder project list"})})," command."]}),"\n",(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{children:"+--------------------------------------+-------------------+\n|                  ID                  |        NAME       |\n+--------------------------------------+-------------------+\n| 086df3e2-f1bb-4b3a-b2fe-0ebd147e1538 | my_minder_project |\n| f9f4aef0-74af-4909-a0c3-0e8ac7fbc38d | another_project   |\n+--------------------------------------+-------------------+\n"})}),"\n",(0,r.jsxs)(n.p,{children:["In this example, the ",(0,r.jsx)(n.code,{children:"my_minder_project"})," project has Project ID\n",(0,r.jsx)(n.code,{children:"086df3e2-f1bb-4b3a-b2fe-0ebd147e1538"}),", and ",(0,r.jsx)(n.code,{children:"another_project"})," has ID\n",(0,r.jsx)(n.code,{children:"f9f4aef0-74af-4909-a0c3-0e8ac7fbc38d"}),". Note that you need to specify the\nproject ",(0,r.jsx)(n.em,{children:"ID"})," when using the ",(0,r.jsx)(n.code,{children:"--project"})," flag or ",(0,r.jsx)(n.code,{children:"MINDER_CONFIG"}),", not the project\nname."]})]})}function h(e={}){const{wrapper:n}={...(0,o.R)(),...e.components};return n?(0,r.jsx)(n,{...e,children:(0,r.jsx)(l,{...e})}):l(e)}},28453:(e,n,i)=>{i.d(n,{R:()=>s,x:()=>c});var t=i(96540);const r={},o=t.createContext(r);function s(e){const n=t.useContext(o);return t.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function c(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(r):e.components||r:s(e.components),t.createElement(o.Provider,{value:n},e.children)}}}]);