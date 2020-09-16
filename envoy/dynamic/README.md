## Background
Refer to current condition, basic balancer for calling graphql is using consul DNS as balancer. It might make the request call unexpected condition, possibly unbalanced traffic.

## Solution
Combine consul with envoy as new modern and reliable balancer. Using dynamic resources configuration, on this approach consul only used for template to update members of envoy cluster.

## Development
Back to root directory, and execute command `make run-dev`, it will spawn envoy and services which will simulating in local machine.

## References
- [Dynamic Configuration](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/operations/dynamic_configuration)
- [Playground](https://www.katacoda.com/envoyproxy/scenarios/file-based-dynamic-routing-configuration)