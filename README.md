<p align="center">
  <img width="40%" src="https://buildbuddy.io/images/header.png"><br/>
  <a href="https://buildbuddy.io">BuildBuddy</a> is an open source Bazel build event viewer<br/><br/>
  <img src="https://img.shields.io/badge/License-MIT-green.svg" />
  <img src="https://img.shields.io/github/workflow/status/tryflame/buildbuddy/CI" />
</p>

# Intro
BuildBuddy is an open source Bazel build event viewer. It helps you collect, view, share and debug build events in a user-friendly web UI.

It's written in Golang and React and can be deployed as a Docker image. It's run both as a [cloud hosted service](https://buildbuddy.io) and can be deployed to your cloud provider or run on-prem. BuildBuddy's core is open sourced in this repo under the [MIT License](https://github.com/tryflame/buildbuddy/blob/master/LICENSE).

# Features

- **[Build summary & build log](https://buildbuddy.io/preview/1-build_log.png)** - a high level overview of the build including who initiated the build, how long it took, how many targets were affected, etc. The build log makes it easy to share stack traces and errors with teammates which makes collaborative debugging easier.
 
- **[Targets & timing](https://buildbuddy.io/preview/2-targets.png)** - see which targets and tests passed / failed along with timing information so you can debug slow builds and tests.
 
- **[Invocation details](https://buildbuddy.io/preview/3-invocation_details.png)** - see all of the explicit flags, implicit options, and environment variables that affect your build. This is particularly useful when a build is working on one machine but not another - you can compare these and see what's different.
 
- **[Artifacts](https://buildbuddy.io/preview/4-artifacts.png)** - get a quick view of all of the build artifacts that were generated by this invocation so you can easily access them.
 
- **[Raw log](https://buildbuddy.io/preview/5-raw_log.png)** - you can really dig into the details here. This is a complete view of all of the events that get sent up via Bazel's build event protocol. If you find yourself digging in here too much, let us know and we'll surface that info in a nicer UI.

# Get started

Getting started with Buildbuddy is simple and free for personal use. Just add these two lines to your `.bazelrc` file.

**.bazelrc**
```
build --bes_results_url=https://app.buildbuddy.io/invocation/
build --bes_backend=grpc://events.buildbuddy.io:1985
```

This will print a **Buildbuddy url** containing your build results at the beginning and end of every Bazel invocation. You can command click / double click on these to view

**Want more control?** Want to set it up for your team? Get up and running fast with the cloud hosted [BuildBuddy.io](https://buildbuddy.io) service.

If you'd like to host your own instance **on-premises** or in the cloud, check out our [getting started](https://github.com/tryflame/buildbuddy/blob/master/SETUP.md) guide.

# Questions?
If you have any questions, e-mail us at hello@tryflame.com. We’d love to chat!
