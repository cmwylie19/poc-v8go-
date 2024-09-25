# Analysis

- [Analysis](#analysis)
  - [Pros and Cons of Running TypeScript from a Go Binary using V8Go](#pros-and-cons-of-running-typescript-from-a-go-binary-using-v8go)
    - [V8Go Pros:](#v8go-pros)
    - [V8Go Cons](#v8go-cons)
  - [Conclusion](#conclusion)
  - [Outro](#outro)



### Pros and Cons of Running TypeScript from a Go Binary using V8Go

#### V8Go Pros

Running TypeScript in a Go environment allows for tighter integration with Kubernetes Go code, leveraging Go's performance and concurrency model.

- **Kubernetes Compatibility**: The V8Go internals should handle chunking of network requests and other I/O-bound operations, making it easier to integrate with Kubernetes Go code.

- **Single Deployment Unit**: Combining TypeScript and Go into a single binary simplifies deployment and may reduce operational overhead.


- **Potential Performance Gains**: By using Go for performance-critical tasks, there might be potential improvements in execution speed compared to pure TypeScript implementations.


- **Use Go Functions**: You can inject Go functions into the Global TypeScript context and use then, theoretically we could take advantage of Go performance and concurrency model and perhaps even use more Go libraries like, for a random example, the CEL framework.

#### V8Go Cons

- **Loss of TypeScript Functionality**: Running TypeScript in V8Go means losing access to many Node.js APIs and features (like the Event Loop), requiring significant effort to reimplement these in Go. This could introduce shaky features in things like Crypto and setTimeout, setInteral, logging and other native Node.js features where we would have to reimplement them in Go.

- **Complex Refactoring Required**: Transitioning to V8Go would necessitate a major refactor of the existing codebase, impacting timelines and stability during the migration process. The timeline would be somewhat undefined if we still want the module author to be able to harness the power of TypeScript as right now they are able to use the whole language, we would have to reimplement a lot of the Node.js features in Go.


- **Runtime Environment Limitations**: Users would not be able to use standard TypeScript/Node.js features like setTimeout, setInterval, and native console logging, which could significantly impact the usability of the framework.

- **Potential Stability Issues**: Introducing Go-based implementations of JavaScript features could lead to stability issues, as the behavior might not match user expectations derived from the original TypeScript/Node.js context.

- **Incompatibility with Other Runtimes**: Refactoring to V8Go would preclude easy transitions to other modern JavaScript runtimes like Deno or Bun in the future, which might offer better performance and features.

- **Increased Maintenance Overhead**: The need to maintain additional Go code for features originally provided by TypeScript/Node.js could lead to increased complexity and maintenance overhead.

- **Development Time and Costs**: The time and resources required for the refactor are uncertain and could lead to budget overruns or delayed timelines, impacting project delivery.

- **Limited Ecosystem**: The rich ecosystem of npm libraries may not be fully leveraged, as many would depend on Node.js functionality that would need to be rewritten in Go.

- **Dev Challenges**: The change in runtime environment will complicate dev strategies, for instance `npx pepr dev` would essentially be a completely different runtime environment than what it is in the cluster.

## Conclusion
While running TypeScript from a Go binary using V8Go presents some potential benefits in terms of integration and performance, the significant drawbacks related to functionality loss, complexity of refactoring, and risks to stability. It may be more advantageous to explore lightweight Go service (K8s Informer) that maintain TypeScript's runtime benefits without introducing the extensive overhead associated with V8Go. The K8s Informer does not affect dev or test, it was implemented and working in just 2 days. We have major expertise in Kubernetes and Go, and the K8s Informer is a natural fit for our existing skill set and allows us to potentially transition to Deno2 or Bun in the future. The idea is maintain the surface area and only change the tiny part that must be changes, the watcher controller and let TypeScript continue to do all the things that it is good at. 




Quick Note on Sobek:
> Although it is maintained by a trusted company, It still suffers from many of the same cons as V8go. This is an exert from their docs, "If most of the work is done in javascript (for example crypto or any other heavy calculations) you are definitely better off with V8", [source](https://github.com/grafana/sobek?tab=readme-ov-file#why-would-i-want-to-use-it-over-a-v8-wrapper).


### Outro

The promise that Pepr gives to Developers is to harness the power of power of a complete programming language to react to changes in the cluster. TypeScript is a full featured, expressive, storngly typed and flexible, however, we would loose that power should be switch to running in a Go binary as a side-effect to changing the Runtime Environment.

The Runtime Environment (like NodeJS or Chrome) is responsible for providing the functionality of Event Loop. As we know that JS engines are capable of executing only one task at a time (or single-threaded), there has to be some different entity that manages providing the JS engine with the next task to execute.
