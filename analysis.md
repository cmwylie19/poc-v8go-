# Analysis

- [Analysis](#analysis)
  - [Decision](#decision)
  - [Pros and Cons of Running JavaScript via V8Go](#pros-and-cons-of-running-javascript-via-v8go)
    - [V8Go Pros](#v8go-pros)
    - [V8Go Cons](#v8go-cons)

## Decision

The Pepr Team preference is not to use V8go for the following reasons:
- v8go is not a drop-in replacement, it does not offer the Runtime Environment of Node.js and we would have to backfill the missing pieces in Go.
- Since we have no way to know what the module author would want to use in the npm/node ecosystem, we would have to backfill everything. (Since missing features could cause instability or we could have bugs in our Go implementation)
- Support for the native Node.js features like setTimeout, setInterval, and native console logging would have to be reimplemented in Go for the Store
- v8go is less supported in comparison to native Kubernetes sigsmachiner go libraries.

Our chosen alternative would be something like the [Watch Informer](https://github.com/cmwylie19/watch-informer) which would run as a sidecar to the controller and would be responsible for watching the resources in the cluster. This would allow us to:
- cut down network pressure by not having to list all the time (the watcher is basically a polling mechanism at this point)
- Move closer to native Kubernetes by using sig-machinery libs which is as supported as Kubernetes itself.
- `npx pepr dev`, `deploy`, `monitor` and ability to switch to deno2 or bun in the future would still be in scope, and the api and surface area of Pepr would remain the same
- Would not have to implement the native Node.js features in Go.
- Less costly/risky to develop and maintain allow us to solve problems in UDS faster.


While running JavaScript from a Go binary using V8Go presents some potential benefits in terms of integration and performance, there are significant drawbacks related to functionality loss, complexity of refactoring, and risks to stability. It is more advantageous to explore lightweight Go service (Watch Informer) that maintain JavaScript's runtime benefits without introducing the extensive overhead associated with V8Go. The Watch Informer does not affect dev or test, it was implemented and working in just 2 days. It is a natural fit for our existing skill set. The Watch Informer solution allows us to maintain the existing surface area and only change the tiny part that must be changed.

### Pros and Cons of Running JavaScript via V8Go

#### V8Go Pros

Running JavaScript in a Go environment allows for tighter integration with Kubernetes Go code, leveraging Go's performance and concurrency model.

- **Kubernetes Compatibility**: The V8Go internals should handle chunking of network requests and other I/O-bound operations, making it easier to integrate with Kubernetes Go code.

- **Single Deployment Unit**: Combining JavaScript and Go into a single binary simplifies deployment and may reduce operational overhead.

- **Potential Performance Gains**: By using Go for performance-critical tasks, there might be potential improvements in execution speed compared to pure JavaScript implementations.

- **Use Go Functions**: You can inject Go functions into the Global JavaScript context and use then, theoretically we could take advantage of Go performance and concurrency model and perhaps even use more Go libraries like, for a random example, the CEL framework.

#### V8Go Cons

- **Loss of JavaScript Functionality**: Running JavaScript in V8Go means losing access to many ready made Node.js APIs, requiring significant effort to reimplement these in Go. This could introduce instability in things like Crypto and setTimeout, setInteral, logging and other native features.

- **Complex Refactoring Required**: Transitioning to V8Go would necessitate a major refactor, impacting timelines and stability during the migration process. The timeline would be somewhat undefined if we still want the module author to be able to harness the power of JavaScript as right now they are able to use the whole language, we would have to reimplement a lot of the Node.js features in Go.

- **Runtime Environment Limitations**: Users would not be able to use many native JavaScript/Node.js features without effort on our part, like native console logging, timeouts, intervals, etc, which could significantly impact the usability of the framework.

- **Potential Stability Issues**: We have to track changes to the Node ecosystem and keep our custom Go implementations up to date with new features. Not tracking the environment fast enough could lead to bugs and instability.

- **Incompatibility with Other Runtimes**: Refactoring to V8Go would preclude easy transitions to other modern JavaScript runtimes like Deno2 or Bun in the future, which might offer better performance and features.

- **Increased Maintenance Overhead**: The need to maintain additional Go code for features originally provided by JavaScript/Node.js could lead to increased complexity and maintenance overhead.

- **Development Time and Costs**: The time and resources required for the refactor are uncertain and could delay timelines on feature work, impacting project delivery.

- **Limited Ecosystem**: The rich ecosystem of npm libraries may not be fully leveragable, as many would depend on Node.js functionality that would need to be rewritten in Go.

- **Dev Challenges**: The change in runtime environment will complicate dev strategies, for instance `npx pepr dev` would be a different runtime environment (Node.js) than what it is in the cluster (V8go).

Quick Note on Sobek (fork of v8go):
> Although it is maintained by a trusted company, It still suffers from many of the same cons as V8go. This is an exert from their docs, "If most of the work is done in javascript (for example crypto or any other heavy calculations) you are definitely better off with V8", [source](https://github.com/grafana/sobek?tab=readme-ov-file#why-would-i-want-to-use-it-over-a-v8-wrapper).

