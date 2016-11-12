# Hob (unstable)

Hob is a self hosted automation manager which facilitates listening for
generic webhook events and propagating them to configured actions and events.

The exact Developer and end-user experience *(api or otherwise)* is still
very much in Flux, so feel free to open any desired Github issues to influence
the design.

With that said, the first versions mainly focus on overall API design and will
not have much *(if any)* ability to be configured via the Web UI.

## API Design

Hob is implemented as a series of events and actions tied together by
a configuration. Events are triggered by sending an HTTP request to an
endpoint *(example: /events/restart_machine)*. Once triggered, the event will
call a single action and return that actions response.

Actions can be chained together as desired by the actions `SyncActions` and
`AsyncActions` *(not yet implemented)*, and custom implementations of the
Action trait can be provided to handle multi-action triggering and responses
differently, as needed.

Both Events and Actions will *(as development progresses)* be flexible to
support multiple methods of communication for simple and advanced event
triggers and action handlers. Eg, to handle the above `restart_machine` event,
a subprocess could be called to simply restart the machine. Alternately, web
server could listen for a webhook and restart the machine when the webhook is
called. For a simpler option, a bash script can loop a curl request on
*/actions/restart_machine/longpoll* on a loop.

Hobs goal is to be simple and flexible for one job: calling actions with
flexibility. No complicated RPC framework is needed.
