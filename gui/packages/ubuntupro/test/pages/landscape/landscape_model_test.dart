import 'package:agentapi/agentapi.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:grpc/grpc.dart';
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';
import 'package:ubuntupro/core/agent_api_client.dart';
import 'package:ubuntupro/pages/landscape/landscape_model.dart';

import 'landscape_model_test.mocks.dart';

@GenerateMocks([AgentApiClient])
void main() {
  group('state management', () {
    final client = MockAgentApiClient();
    test('active', () {
      final model = LandscapeModel(client);
      expect(model.configType, LandscapeConfigType.selfHosted);

      for (final type in LandscapeConfigType.values) {
        model.setConfigType(type);
        expect(model.configType, type);
      }
    });

    test('notify changes', () {
      // Verifies that all set* methods notify listeners.

      final model = LandscapeModel(client);
      var notified = false;
      model.addListener(() {
        notified = true;
      });

      model.setConfigType(LandscapeConfigType.custom);
      expect(notified, isTrue);
      notified = false;

      model.setCustomConfigPath(customConf);
      expect(notified, isTrue);
      notified = false;

      model.setConfigType(LandscapeConfigType.saas);
      expect(notified, isTrue);
      notified = false;

      model.setAccountName('testuser');
      expect(notified, isTrue);
      notified = false;

      model.setSaasRegistrationKey('123');
      expect(notified, isTrue);
      notified = false;

      model.setConfigType(LandscapeConfigType.selfHosted);
      expect(notified, isTrue);
      notified = false;

      model.setFqdn(testFqdn);
      expect(notified, isTrue);
      notified = false;

      model.setSelfHostedRegistrationKey('123');
      expect(notified, isTrue);
      notified = false;

      model.setSslKeyPath(customConf);
      expect(notified, isTrue);
      notified = false;
    });

    test('assertions', () {
      // Verifies that methods throw assertions when called under non-relevant scenarios.

      final model = LandscapeModel(client);

      model.setConfigType(LandscapeConfigType.saas);
      // Those assertions exist because the methods are not relevant for the current config type.
      // Allowing those conditions to proceed could contribute to hide logic errors.
      expect(() => model.setCustomConfigPath(customConf), throwsAssertionError);
      expect(() => model.setFqdn(testFqdn), throwsAssertionError);
      expect(() => model.setSslKeyPath(customConf), throwsAssertionError);
      expect(
        () => model.setSelfHostedRegistrationKey('123'),
        throwsAssertionError,
      );
      expect(() => model.setCustomConfigPath(customConf), throwsAssertionError);

      model.setConfigType(LandscapeConfigType.selfHosted);
      expect(() => model.setAccountName('testuser'), throwsAssertionError);
      expect(() => model.setSaasRegistrationKey('123'), throwsAssertionError);
      expect(() => model.setCustomConfigPath(customConf), throwsAssertionError);

      model.setConfigType(LandscapeConfigType.custom);
      expect(() => model.setSslKeyPath(customConf), throwsAssertionError);
      expect(() => model.setAccountName('testuser'), throwsAssertionError);
      expect(() => model.setSaasRegistrationKey('123'), throwsAssertionError);
      expect(() => model.setFqdn(testFqdn), throwsAssertionError);
      expect(
        () => model.setSelfHostedRegistrationKey('123'),
        throwsAssertionError,
      );
      expect(() => model.setSslKeyPath(customConf), throwsAssertionError);
    });
  });

  group('apply config', () {
    const msg = 'test message';
    const error = GrpcError.custom(StatusCode.unavailable, msg);
    test('saas', () async {
      final client = MockAgentApiClient();
      when(client.applyLandscapeConfig(any)).thenAnswer(
        (_) async => throw error,
      );
      final model = LandscapeModel(client);

      model.setConfigType(LandscapeConfigType.saas);
      expect(model.applyConfig, throwsAssertionError);

      model.setAccountName('testaccount');
      var err = await model.applyConfig();
      expect(err, msg);

      when(client.applyLandscapeConfig(any)).thenAnswer(
        (_) async => LandscapeSource()..ensureUser(),
      );

      model.setAccountName('testaccount');
      err = await model.applyConfig();
      expect(err, isNull);
    });
    test('self-hosted', () async {
      final client = MockAgentApiClient();
      when(client.applyLandscapeConfig(any)).thenAnswer(
        (_) async => throw error,
      );
      final model = LandscapeModel(client);

      model.setConfigType(LandscapeConfigType.selfHosted);
      expect(model.applyConfig, throwsAssertionError);

      model.setFqdn(testFqdn);
      var err = await model.applyConfig();
      expect(err, msg);

      when(client.applyLandscapeConfig(any)).thenAnswer(
        (_) async => LandscapeSource()..ensureUser(),
      );

      model.setFqdn(testFqdn);
      err = await model.applyConfig();
      expect(err, isNull);

      model.setSslKeyPath(caCert);
      err = await model.applyConfig();
      expect(err, isNull);
    });
    test('custom', () async {
      final client = MockAgentApiClient();
      when(client.applyLandscapeConfig(any)).thenAnswer(
        (_) async => throw error,
      );
      final model = LandscapeModel(client);

      model.setConfigType(LandscapeConfigType.custom);
      expect(model.applyConfig, throwsAssertionError);

      model.setCustomConfigPath(customConf);
      var err = await model.applyConfig();
      expect(err, msg);

      when(client.applyLandscapeConfig(any)).thenAnswer(
        (_) async => LandscapeSource()..ensureUser(),
      );

      model.setCustomConfigPath(customConf);
      err = await model.applyConfig();
      expect(err, isNull);
    });
  });
}

const customConf = './test/testdata/landscape/custom.conf';
const saasURL = 'https://landscape.canonical.com';
const testFqdn = 'test.landscape.company.com';
const caCert = './test/testdata/certs/ca_cert.pem';
