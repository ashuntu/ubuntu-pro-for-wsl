// Mocks generated by Mockito 5.4.4 from annotations
// in ubuntupro/test/startup/startup_page_test.dart.
// Do not manually edit this file.

// ignore_for_file: no_leading_underscores_for_library_prefixes
import 'dart:async' as _i5;

import 'package:agentapi/agentapi.dart' as _i3;
import 'package:mockito/mockito.dart' as _i1;
import 'package:ubuntupro/core/agent_api_client.dart' as _i2;
import 'package:ubuntupro/pages/startup/agent_monitor.dart' as _i4;

// ignore_for_file: type=lint
// ignore_for_file: avoid_redundant_argument_values
// ignore_for_file: avoid_setters_without_getters
// ignore_for_file: comment_references
// ignore_for_file: deprecated_member_use
// ignore_for_file: deprecated_member_use_from_same_package
// ignore_for_file: implementation_imports
// ignore_for_file: invalid_use_of_visible_for_testing_member
// ignore_for_file: prefer_const_constructors
// ignore_for_file: unnecessary_parenthesis
// ignore_for_file: camel_case_types
// ignore_for_file: subtype_of_sealed_class

class _FakeAgentApiClient_0 extends _i1.SmartFake
    implements _i2.AgentApiClient {
  _FakeAgentApiClient_0(
    Object parent,
    Invocation parentInvocation,
  ) : super(
          parent,
          parentInvocation,
        );
}

class _FakeSubscriptionInfo_1 extends _i1.SmartFake
    implements _i3.SubscriptionInfo {
  _FakeSubscriptionInfo_1(
    Object parent,
    Invocation parentInvocation,
  ) : super(
          parent,
          parentInvocation,
        );
}

/// A class which mocks [AgentStartupMonitor].
///
/// See the documentation for Mockito's code generation for more information.
class MockAgentStartupMonitor extends _i1.Mock
    implements _i4.AgentStartupMonitor {
  MockAgentStartupMonitor() {
    _i1.throwOnMissingStub(this);
  }

  @override
  _i4.AgentLauncher get agentLauncher => (super.noSuchMethod(
        Invocation.getter(#agentLauncher),
        returnValue: () => _i5.Future<bool>.value(false),
      ) as _i4.AgentLauncher);

  @override
  _i4.ApiClientFactory get clientFactory => (super.noSuchMethod(
        Invocation.getter(#clientFactory),
        returnValue: (int port) => _FakeAgentApiClient_0(
          this,
          Invocation.getter(#clientFactory),
        ),
      ) as _i4.ApiClientFactory);

  @override
  _i4.AgentApiCallback get onClient => (super.noSuchMethod(
        Invocation.getter(#onClient),
        returnValue: (_i2.AgentApiClient __p0) => null,
      ) as _i4.AgentApiCallback);

  @override
  _i5.Stream<_i4.AgentState> start({
    Duration? interval = const Duration(seconds: 1),
    Duration? timeout = const Duration(seconds: 5),
  }) =>
      (super.noSuchMethod(
        Invocation.method(
          #start,
          [],
          {
            #interval: interval,
            #timeout: timeout,
          },
        ),
        returnValue: _i5.Stream<_i4.AgentState>.empty(),
      ) as _i5.Stream<_i4.AgentState>);

  @override
  _i5.Future<void> reset() => (super.noSuchMethod(
        Invocation.method(
          #reset,
          [],
        ),
        returnValue: _i5.Future<void>.value(),
        returnValueForMissingStub: _i5.Future<void>.value(),
      ) as _i5.Future<void>);
}

/// A class which mocks [AgentApiClient].
///
/// See the documentation for Mockito's code generation for more information.
class MockAgentApiClient extends _i1.Mock implements _i2.AgentApiClient {
  MockAgentApiClient() {
    _i1.throwOnMissingStub(this);
  }

  @override
  _i5.Future<_i3.SubscriptionInfo> applyProToken(String? token) =>
      (super.noSuchMethod(
        Invocation.method(
          #applyProToken,
          [token],
        ),
        returnValue:
            _i5.Future<_i3.SubscriptionInfo>.value(_FakeSubscriptionInfo_1(
          this,
          Invocation.method(
            #applyProToken,
            [token],
          ),
        )),
      ) as _i5.Future<_i3.SubscriptionInfo>);

  @override
  _i5.Future<bool> ping() => (super.noSuchMethod(
        Invocation.method(
          #ping,
          [],
        ),
        returnValue: _i5.Future<bool>.value(false),
      ) as _i5.Future<bool>);

  @override
  _i5.Future<_i3.SubscriptionInfo> subscriptionInfo() => (super.noSuchMethod(
        Invocation.method(
          #subscriptionInfo,
          [],
        ),
        returnValue:
            _i5.Future<_i3.SubscriptionInfo>.value(_FakeSubscriptionInfo_1(
          this,
          Invocation.method(
            #subscriptionInfo,
            [],
          ),
        )),
      ) as _i5.Future<_i3.SubscriptionInfo>);

  @override
  _i5.Future<_i3.SubscriptionInfo> notifyPurchase() => (super.noSuchMethod(
        Invocation.method(
          #notifyPurchase,
          [],
        ),
        returnValue:
            _i5.Future<_i3.SubscriptionInfo>.value(_FakeSubscriptionInfo_1(
          this,
          Invocation.method(
            #notifyPurchase,
            [],
          ),
        )),
      ) as _i5.Future<_i3.SubscriptionInfo>);
}
