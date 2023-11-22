**Thought Processs** <br />
The whole idea is to create a remote config that can be used to
store application configurations in a very simple way.
It will be broken down into versions.<br />

Versions <br />
1. Remote Configuration, store and retrieve config data <br />
2. We will implement Feature Flag (called flagger disabled by default)
3. We will implement secret (and integrate with kubernetes and few others)

<u> Create account 3 simple steps </u> <br />
<ol>
<li> Input your email, password and workspace name</li>
<li>Validate Your Email </li>
<li>Create a Project </li>
</ol>

<hr />

<h1> HOW TO </h1>
1. Each account has account id,  secret id and lists of workspace <br />
2. Each workspace has an id and state (ACTIVE, INACTIVE) <br />
3. Each workspace has lists of projects <br />
4. Each Project in turn has several environments. By default 3 (dev, pilot, prod) environments will be created
for each project created automatically. <br />
5. Each Project has by two state (ACTIVE, INACTIVE), same with environments <br />
6. Each Project has settings (Project Settings) <br />
7. Each Workspace has settings <br />
8. Premium Account can invite other user (known as members) <br />
9. Each Account has some set of Permissions <br />
10. Default output format is json but you have the options of choosing default one or in each project from the following (YAML, TOML, JSON) <br />
11. You can import config in any format <br>
12. The system has CLI, SDK and web interface for interacting with the api <br>
13. By Default, The API is not directly accessible via rest except enabled in account settings <br>
14. Will implement with few providers (terraform, cloud formation) etc





