// Code generated by protoc-gen-go-j5. DO NOT EDIT.

package github_pb

import (
	context "context"
	fmt "fmt"

	psm_j5pb "github.com/pentops/j5/gen/j5/state/v1/psm_j5pb"
	psm "github.com/pentops/j5/lib/psm"
	sqrlx "github.com/pentops/sqrlx.go/sqrlx"
)

// PSM RepoPSM

type RepoPSM = psm.StateMachine[
	*RepoKeys,      // implements psm.IKeyset
	*RepoState,     // implements psm.IState
	RepoStatus,     // implements psm.IStatusEnum
	*RepoStateData, // implements psm.IStateData
	*RepoEvent,     // implements psm.IEvent
	RepoPSMEvent,   // implements psm.IInnerEvent
]

type RepoPSMDB = psm.DBStateMachine[
	*RepoKeys,      // implements psm.IKeyset
	*RepoState,     // implements psm.IState
	RepoStatus,     // implements psm.IStatusEnum
	*RepoStateData, // implements psm.IStateData
	*RepoEvent,     // implements psm.IEvent
	RepoPSMEvent,   // implements psm.IInnerEvent
]

type RepoPSMEventSpec = psm.EventSpec[
	*RepoKeys,      // implements psm.IKeyset
	*RepoState,     // implements psm.IState
	RepoStatus,     // implements psm.IStatusEnum
	*RepoStateData, // implements psm.IStateData
	*RepoEvent,     // implements psm.IEvent
	RepoPSMEvent,   // implements psm.IInnerEvent
]

type RepoPSMHookBaton = psm.HookBaton[
	*RepoKeys,      // implements psm.IKeyset
	*RepoState,     // implements psm.IState
	RepoStatus,     // implements psm.IStatusEnum
	*RepoStateData, // implements psm.IStateData
	*RepoEvent,     // implements psm.IEvent
	RepoPSMEvent,   // implements psm.IInnerEvent
]

type RepoPSMFullBaton = psm.CallbackBaton[
	*RepoKeys,      // implements psm.IKeyset
	*RepoState,     // implements psm.IState
	RepoStatus,     // implements psm.IStatusEnum
	*RepoStateData, // implements psm.IStateData
	*RepoEvent,     // implements psm.IEvent
	RepoPSMEvent,   // implements psm.IInnerEvent
]

type RepoPSMEventKey = string

const (
	RepoPSMEventNil             RepoPSMEventKey = "<nil>"
	RepoPSMEventConfigure       RepoPSMEventKey = "configure"
	RepoPSMEventConfigureBranch RepoPSMEventKey = "configure_branch"
	RepoPSMEventRemoveBranch    RepoPSMEventKey = "remove_branch"
)

// EXTEND RepoKeys with the psm.IKeyset interface

// PSMIsSet is a helper for != nil, which does not work with generic parameters
func (msg *RepoKeys) PSMIsSet() bool {
	return msg != nil
}

// PSMFullName returns the full name of state machine with package prefix
func (msg *RepoKeys) PSMFullName() string {
	return "j5.builds.github.v1.repo"
}
func (msg *RepoKeys) PSMKeyValues() (map[string]any, error) {
	keyset := map[string]any{
		"owner": msg.Owner,
		"name":  msg.Name,
	}
	return keyset, nil
}

// EXTEND RepoState with the psm.IState interface

// PSMIsSet is a helper for != nil, which does not work with generic parameters
func (msg *RepoState) PSMIsSet() bool {
	return msg != nil
}

func (msg *RepoState) PSMMetadata() *psm_j5pb.StateMetadata {
	if msg.Metadata == nil {
		msg.Metadata = &psm_j5pb.StateMetadata{}
	}
	return msg.Metadata
}

func (msg *RepoState) PSMKeys() *RepoKeys {
	return msg.Keys
}

func (msg *RepoState) SetStatus(status RepoStatus) {
	msg.Status = status
}

func (msg *RepoState) SetPSMKeys(inner *RepoKeys) {
	msg.Keys = inner
}

func (msg *RepoState) PSMData() *RepoStateData {
	if msg.Data == nil {
		msg.Data = &RepoStateData{}
	}
	return msg.Data
}

// EXTEND RepoStateData with the psm.IStateData interface

// PSMIsSet is a helper for != nil, which does not work with generic parameters
func (msg *RepoStateData) PSMIsSet() bool {
	return msg != nil
}

// EXTEND RepoEvent with the psm.IEvent interface

// PSMIsSet is a helper for != nil, which does not work with generic parameters
func (msg *RepoEvent) PSMIsSet() bool {
	return msg != nil
}

func (msg *RepoEvent) PSMMetadata() *psm_j5pb.EventMetadata {
	if msg.Metadata == nil {
		msg.Metadata = &psm_j5pb.EventMetadata{}
	}
	return msg.Metadata
}

func (msg *RepoEvent) PSMKeys() *RepoKeys {
	return msg.Keys
}

func (msg *RepoEvent) SetPSMKeys(inner *RepoKeys) {
	msg.Keys = inner
}

// PSMEventKey returns the RepoPSMEventPSMEventKey for the event, implementing psm.IEvent
func (msg *RepoEvent) PSMEventKey() RepoPSMEventKey {
	tt := msg.UnwrapPSMEvent()
	if tt == nil {
		return RepoPSMEventNil
	}
	return tt.PSMEventKey()
}

// UnwrapPSMEvent implements psm.IEvent, returning the inner event message
func (msg *RepoEvent) UnwrapPSMEvent() RepoPSMEvent {
	if msg == nil {
		return nil
	}
	if msg.Event == nil {
		return nil
	}
	switch v := msg.Event.Type.(type) {
	case *RepoEventType_Configure_:
		return v.Configure
	case *RepoEventType_ConfigureBranch_:
		return v.ConfigureBranch
	case *RepoEventType_RemoveBranch_:
		return v.RemoveBranch
	default:
		return nil
	}
}

// SetPSMEvent sets the inner event message from a concrete type, implementing psm.IEvent
func (msg *RepoEvent) SetPSMEvent(inner RepoPSMEvent) error {
	if msg.Event == nil {
		msg.Event = &RepoEventType{}
	}
	switch v := inner.(type) {
	case *RepoEventType_Configure:
		msg.Event.Type = &RepoEventType_Configure_{Configure: v}
	case *RepoEventType_ConfigureBranch:
		msg.Event.Type = &RepoEventType_ConfigureBranch_{ConfigureBranch: v}
	case *RepoEventType_RemoveBranch:
		msg.Event.Type = &RepoEventType_RemoveBranch_{RemoveBranch: v}
	default:
		return fmt.Errorf("invalid type %T for RepoEventType", v)
	}
	return nil
}

type RepoPSMEvent interface {
	psm.IInnerEvent
	PSMEventKey() RepoPSMEventKey
}

// EXTEND RepoEventType_Configure with the RepoPSMEvent interface

// PSMIsSet is a helper for != nil, which does not work with generic parameters
func (msg *RepoEventType_Configure) PSMIsSet() bool {
	return msg != nil
}

func (*RepoEventType_Configure) PSMEventKey() RepoPSMEventKey {
	return RepoPSMEventConfigure
}

// EXTEND RepoEventType_ConfigureBranch with the RepoPSMEvent interface

// PSMIsSet is a helper for != nil, which does not work with generic parameters
func (msg *RepoEventType_ConfigureBranch) PSMIsSet() bool {
	return msg != nil
}

func (*RepoEventType_ConfigureBranch) PSMEventKey() RepoPSMEventKey {
	return RepoPSMEventConfigureBranch
}

// EXTEND RepoEventType_RemoveBranch with the RepoPSMEvent interface

// PSMIsSet is a helper for != nil, which does not work with generic parameters
func (msg *RepoEventType_RemoveBranch) PSMIsSet() bool {
	return msg != nil
}

func (*RepoEventType_RemoveBranch) PSMEventKey() RepoPSMEventKey {
	return RepoPSMEventRemoveBranch
}

func RepoPSMBuilder() *psm.StateMachineConfig[
	*RepoKeys,      // implements psm.IKeyset
	*RepoState,     // implements psm.IState
	RepoStatus,     // implements psm.IStatusEnum
	*RepoStateData, // implements psm.IStateData
	*RepoEvent,     // implements psm.IEvent
	RepoPSMEvent,   // implements psm.IInnerEvent
] {
	return &psm.StateMachineConfig[
		*RepoKeys,      // implements psm.IKeyset
		*RepoState,     // implements psm.IState
		RepoStatus,     // implements psm.IStatusEnum
		*RepoStateData, // implements psm.IStateData
		*RepoEvent,     // implements psm.IEvent
		RepoPSMEvent,   // implements psm.IInnerEvent
	]{}
}

// RepoPSMMutation runs at the start of a transition to merge the event information into the state data object. The state object is mutable in this context.
func RepoPSMMutation[SE RepoPSMEvent](cb func(*RepoStateData, SE) error) psm.TransitionMutation[
	*RepoKeys,      // implements psm.IKeyset
	*RepoState,     // implements psm.IState
	RepoStatus,     // implements psm.IStatusEnum
	*RepoStateData, // implements psm.IStateData
	*RepoEvent,     // implements psm.IEvent
	RepoPSMEvent,   // implements psm.IInnerEvent
	SE,             // Specific event type for the transition
] {
	return psm.TransitionMutation[
		*RepoKeys,      // implements psm.IKeyset
		*RepoState,     // implements psm.IState
		RepoStatus,     // implements psm.IStatusEnum
		*RepoStateData, // implements psm.IStateData
		*RepoEvent,     // implements psm.IEvent
		RepoPSMEvent,   // implements psm.IInnerEvent
		SE,             // Specific event type for the transition
	](cb)
}

// RepoPSMLogicHook runs after the mutation is complete. This hook can trigger side effects, including chained events, which are additional events processed by the state machine. Use this for Business Logic which determines the 'next step' in processing.
func RepoPSMLogicHook[
	SE RepoPSMEvent,
](
	cb func(
		context.Context,
		RepoPSMHookBaton,
		*RepoState,
		SE,
	) error) psm.TransitionHook[
	*RepoKeys,      // implements psm.IKeyset
	*RepoState,     // implements psm.IState
	RepoStatus,     // implements psm.IStatusEnum
	*RepoStateData, // implements psm.IStateData
	*RepoEvent,     // implements psm.IEvent
	RepoPSMEvent,   // implements psm.IInnerEvent
] {
	eventType := (*new(SE)).PSMEventKey()
	return psm.TransitionHook[
		*RepoKeys,      // implements psm.IKeyset
		*RepoState,     // implements psm.IState
		RepoStatus,     // implements psm.IStatusEnum
		*RepoStateData, // implements psm.IStateData
		*RepoEvent,     // implements psm.IEvent
		RepoPSMEvent,   // implements psm.IInnerEvent
	]{
		Callback: func(ctx context.Context, tx sqrlx.Transaction, baton RepoPSMFullBaton, state *RepoState, event *RepoEvent) error {
			asType, ok := any(event.UnwrapPSMEvent()).(SE)
			if !ok {
				name := event.ProtoReflect().Descriptor().FullName()
				return fmt.Errorf("unexpected event type in transition: %s [IE] does not match [SE] (%T)", name, new(SE))
			}
			return cb(ctx, baton, state, asType)
		},
		EventType:   eventType,
		RunOnFollow: false,
	}
}

// RepoPSMDataHook runs after the mutations, and can be used to update data in tables which are not controlled as the state machine, e.g. for pre-calculating fields for performance reasons. Use of this hook prevents (future) transaction optimizations, as the transaction state when the function is called must needs to match the processing state, but only for this single transition, unlike the GeneralEventDataHook.
func RepoPSMDataHook[
	SE RepoPSMEvent,
](
	cb func(
		context.Context,
		sqrlx.Transaction,
		*RepoState,
		SE,
	) error) psm.TransitionHook[
	*RepoKeys,      // implements psm.IKeyset
	*RepoState,     // implements psm.IState
	RepoStatus,     // implements psm.IStatusEnum
	*RepoStateData, // implements psm.IStateData
	*RepoEvent,     // implements psm.IEvent
	RepoPSMEvent,   // implements psm.IInnerEvent
] {
	eventType := (*new(SE)).PSMEventKey()
	return psm.TransitionHook[
		*RepoKeys,      // implements psm.IKeyset
		*RepoState,     // implements psm.IState
		RepoStatus,     // implements psm.IStatusEnum
		*RepoStateData, // implements psm.IStateData
		*RepoEvent,     // implements psm.IEvent
		RepoPSMEvent,   // implements psm.IInnerEvent
	]{
		Callback: func(ctx context.Context, tx sqrlx.Transaction, baton RepoPSMFullBaton, state *RepoState, event *RepoEvent) error {
			asType, ok := any(event.UnwrapPSMEvent()).(SE)
			if !ok {
				name := event.ProtoReflect().Descriptor().FullName()
				return fmt.Errorf("unexpected event type in transition: %s [IE] does not match [SE] (%T)", name, new(SE))
			}
			return cb(ctx, tx, state, asType)
		},
		EventType:   eventType,
		RunOnFollow: true,
	}
}

// RepoPSMLinkHook runs after the mutation and logic hook, and can be used to link the state machine to other state machines in the same database transaction
func RepoPSMLinkHook[
	SE RepoPSMEvent,
	DK psm.IKeyset,
	DIE psm.IInnerEvent,
](
	linkDestination psm.LinkDestination[DK, DIE],
	cb func(
		context.Context,
		*RepoState,
		SE,
		func(DK, DIE),
	) error) psm.TransitionHook[
	*RepoKeys,      // implements psm.IKeyset
	*RepoState,     // implements psm.IState
	RepoStatus,     // implements psm.IStatusEnum
	*RepoStateData, // implements psm.IStateData
	*RepoEvent,     // implements psm.IEvent
	RepoPSMEvent,   // implements psm.IInnerEvent
] {
	eventType := (*new(SE)).PSMEventKey()
	wrapped := func(ctx context.Context, tx sqrlx.Transaction, state *RepoState, event SE, add func(DK, DIE)) error {
		return cb(ctx, state, event, add)
	}
	return psm.TransitionHook[
		*RepoKeys,      // implements psm.IKeyset
		*RepoState,     // implements psm.IState
		RepoStatus,     // implements psm.IStatusEnum
		*RepoStateData, // implements psm.IStateData
		*RepoEvent,     // implements psm.IEvent
		RepoPSMEvent,   // implements psm.IInnerEvent
	]{
		Callback: func(ctx context.Context, tx sqrlx.Transaction, baton RepoPSMFullBaton, state *RepoState, event *RepoEvent) error {
			return psm.RunLinkHook(ctx, linkDestination, wrapped, tx, state, event)
		},
		EventType:   eventType,
		RunOnFollow: false,
	}
}

// RepoPSMLinkDBHook like LinkHook, but has access to the current transaction for reads only (not enforced), use in place of controller logic to look up existing state.
func RepoPSMLinkDBHook[
	SE RepoPSMEvent,
	DK psm.IKeyset,
	DIE psm.IInnerEvent,
](
	linkDestination psm.LinkDestination[DK, DIE],
	cb func(
		context.Context,
		sqrlx.Transaction,
		*RepoState,
		SE,
		func(DK, DIE),
	) error) psm.TransitionHook[
	*RepoKeys,      // implements psm.IKeyset
	*RepoState,     // implements psm.IState
	RepoStatus,     // implements psm.IStatusEnum
	*RepoStateData, // implements psm.IStateData
	*RepoEvent,     // implements psm.IEvent
	RepoPSMEvent,   // implements psm.IInnerEvent
] {
	eventType := (*new(SE)).PSMEventKey()
	return psm.TransitionHook[
		*RepoKeys,      // implements psm.IKeyset
		*RepoState,     // implements psm.IState
		RepoStatus,     // implements psm.IStatusEnum
		*RepoStateData, // implements psm.IStateData
		*RepoEvent,     // implements psm.IEvent
		RepoPSMEvent,   // implements psm.IInnerEvent
	]{
		Callback: func(ctx context.Context, tx sqrlx.Transaction, baton RepoPSMFullBaton, state *RepoState, event *RepoEvent) error {
			return psm.RunLinkHook(ctx, linkDestination, cb, tx, state, event)
		},
		EventType:   eventType,
		RunOnFollow: false,
	}
}

// RepoPSMGeneralLogicHook runs once per transition at the state-machine level regardless of which transition / event is being processed. It runs exactly once per transition, with the state object in the final state after the transition but prior to processing any further events. Chained events are added to the *end* of the event queue for the transaction, and side effects are published (as always) when the transaction is committed. The function MUST be pure, i.e. It MUST NOT produce any side-effects outside of the HookBaton, and MUST NOT modify the state.
func RepoPSMGeneralLogicHook(
	cb func(
		context.Context,
		RepoPSMHookBaton,
		*RepoState,
		*RepoEvent,
	) error) psm.GeneralEventHook[
	*RepoKeys,      // implements psm.IKeyset
	*RepoState,     // implements psm.IState
	RepoStatus,     // implements psm.IStatusEnum
	*RepoStateData, // implements psm.IStateData
	*RepoEvent,     // implements psm.IEvent
	RepoPSMEvent,   // implements psm.IInnerEvent
] {
	return psm.GeneralEventHook[
		*RepoKeys,      // implements psm.IKeyset
		*RepoState,     // implements psm.IState
		RepoStatus,     // implements psm.IStatusEnum
		*RepoStateData, // implements psm.IStateData
		*RepoEvent,     // implements psm.IEvent
		RepoPSMEvent,   // implements psm.IInnerEvent
	]{
		Callback: func(
			ctx context.Context,
			tx sqrlx.Transaction,
			baton RepoPSMFullBaton,
			state *RepoState,
			event *RepoEvent,
		) error {
			return cb(ctx, baton, state, event)
		},
		RunOnFollow: false,
	}
}

// RepoPSMGeneralStateDataHook runs at the state-machine level regardless of which transition / event is being processed. It runs at-least once before committing a database transaction after multiple transitions are complete. This hook has access only to the final state after the transitions and is used to update other tables based on the resulting state. It MUST be idempotent, it may be called after injecting externally-held state data.
func RepoPSMGeneralStateDataHook(
	cb func(
		context.Context,
		sqrlx.Transaction,
		*RepoState,
	) error) psm.GeneralStateHook[
	*RepoKeys,      // implements psm.IKeyset
	*RepoState,     // implements psm.IState
	RepoStatus,     // implements psm.IStatusEnum
	*RepoStateData, // implements psm.IStateData
	*RepoEvent,     // implements psm.IEvent
	RepoPSMEvent,   // implements psm.IInnerEvent
] {
	return psm.GeneralStateHook[
		*RepoKeys,      // implements psm.IKeyset
		*RepoState,     // implements psm.IState
		RepoStatus,     // implements psm.IStatusEnum
		*RepoStateData, // implements psm.IStateData
		*RepoEvent,     // implements psm.IEvent
		RepoPSMEvent,   // implements psm.IInnerEvent
	]{
		Callback: func(
			ctx context.Context,
			tx sqrlx.Transaction,
			baton RepoPSMFullBaton,
			state *RepoState,
		) error {
			return cb(ctx, tx, state)
		},
		RunOnFollow: true,
	}
}

// RepoPSMGeneralEventDataHook runs after each transition at the state-machine level regardless of which transition / event is being processed. It runs exactly once per transition, before any other events are processed. The presence of this hook type prevents (future) transaction optimizations, so should be used sparingly.
func RepoPSMGeneralEventDataHook(
	cb func(
		context.Context,
		sqrlx.Transaction,
		*RepoState,
		*RepoEvent,
	) error) psm.GeneralEventHook[
	*RepoKeys,      // implements psm.IKeyset
	*RepoState,     // implements psm.IState
	RepoStatus,     // implements psm.IStatusEnum
	*RepoStateData, // implements psm.IStateData
	*RepoEvent,     // implements psm.IEvent
	RepoPSMEvent,   // implements psm.IInnerEvent
] {
	return psm.GeneralEventHook[
		*RepoKeys,      // implements psm.IKeyset
		*RepoState,     // implements psm.IState
		RepoStatus,     // implements psm.IStatusEnum
		*RepoStateData, // implements psm.IStateData
		*RepoEvent,     // implements psm.IEvent
		RepoPSMEvent,   // implements psm.IInnerEvent
	]{
		Callback: func(
			ctx context.Context,
			tx sqrlx.Transaction,
			baton RepoPSMFullBaton,
			state *RepoState,
			event *RepoEvent,
		) error {
			return cb(ctx, tx, state, event)
		},
		RunOnFollow: true,
	}
}

// RepoPSMEventPublishHook  EventPublishHook runs for each transition, at least once before committing a database transaction after multiple transitions are complete. It should publish a derived version of the event using the publisher.
func RepoPSMEventPublishHook(
	cb func(
		context.Context,
		psm.Publisher,
		*RepoState,
		*RepoEvent,
	) error) psm.GeneralEventHook[
	*RepoKeys,      // implements psm.IKeyset
	*RepoState,     // implements psm.IState
	RepoStatus,     // implements psm.IStatusEnum
	*RepoStateData, // implements psm.IStateData
	*RepoEvent,     // implements psm.IEvent
	RepoPSMEvent,   // implements psm.IInnerEvent
] {
	return psm.GeneralEventHook[
		*RepoKeys,      // implements psm.IKeyset
		*RepoState,     // implements psm.IState
		RepoStatus,     // implements psm.IStatusEnum
		*RepoStateData, // implements psm.IStateData
		*RepoEvent,     // implements psm.IEvent
		RepoPSMEvent,   // implements psm.IInnerEvent
	]{
		Callback: func(
			ctx context.Context,
			tx sqrlx.Transaction,
			baton RepoPSMFullBaton,
			state *RepoState,
			event *RepoEvent,
		) error {
			return cb(ctx, baton, state, event)
		},
		RunOnFollow: false,
	}
}

// RepoPSMUpsertPublishHook runs for each transition, at least once before committing a database transaction after multiple transitions are complete. It should publish a derived version of the event using the publisher.
func RepoPSMUpsertPublishHook(
	cb func(
		context.Context,
		psm.Publisher,
		*RepoState,
	) error) psm.GeneralStateHook[
	*RepoKeys,      // implements psm.IKeyset
	*RepoState,     // implements psm.IState
	RepoStatus,     // implements psm.IStatusEnum
	*RepoStateData, // implements psm.IStateData
	*RepoEvent,     // implements psm.IEvent
	RepoPSMEvent,   // implements psm.IInnerEvent
] {
	return psm.GeneralStateHook[
		*RepoKeys,      // implements psm.IKeyset
		*RepoState,     // implements psm.IState
		RepoStatus,     // implements psm.IStatusEnum
		*RepoStateData, // implements psm.IStateData
		*RepoEvent,     // implements psm.IEvent
		RepoPSMEvent,   // implements psm.IInnerEvent
	]{
		Callback: func(
			ctx context.Context,
			tx sqrlx.Transaction,
			baton RepoPSMFullBaton,
			state *RepoState,
		) error {
			return cb(ctx, baton, state)
		},
		RunOnFollow: false,
	}
}

func (event *RepoEvent) EventPublishMetadata() *psm_j5pb.EventPublishMetadata {
	tenantKeys := make([]*psm_j5pb.EventTenant, 0)
	return &psm_j5pb.EventPublishMetadata{
		EventId:   event.Metadata.EventId,
		Sequence:  event.Metadata.Sequence,
		Timestamp: event.Metadata.Timestamp,
		Cause:     event.Metadata.Cause,
		Auth: &psm_j5pb.PublishAuth{
			TenantKeys: tenantKeys,
		},
	}
}
