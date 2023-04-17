import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import '../../core/agent_api_client.dart';
import '../../core/either_value_notifier.dart';
import '../../core/pro_token.dart';

extension TokenErrorl10n on TokenError {
  /// Allows representing the [TokenError] enum as a String.
  String localize(AppLocalizations lang) {
    switch (this) {
      case TokenError.empty:
        return lang.tokenErrorEmpty;
      case TokenError.tooShort:
        return lang.tokenErrorTooShort;
      case TokenError.tooLong:
        return lang.tokenErrorTooLong;
      case TokenError.invalidPrefix:
        return lang.tokenErrorInvalidPrefix;
      case TokenError.invalidEncoding:
        return lang.tokenErrorInvalidEncoding;
      default:
        throw UnimplementedError(toString());
    }
  }
}

/// The view-model for the [EnterProTokenPage].
/// Since we don't want to start the UI with an error due the text field being
/// empty, this stores a nullable [ProToken] object
class EnterProTokenModel extends EitherValueNotifier<TokenError, ProToken?> {
  EnterProTokenModel(this.client) : super.ok(null);

  final AgentApiClient client;

  String? get token => valueOrNull?.value;

  bool get hasError => value.isLeft;

  void update(String token) {
    value = ProToken.create(token);
  }

  Future<void> apply() async {
    if (value.isRight) {
      return client.applyProToken(valueOrNull!.value);
    }
  }
}

/// Computes the required width to display comfortable a Pro token with maximum
///  length in logical pixels.
double? maxTokenWidth({
  required double fontSize,
  required double textScale,
}) {
  return fontSize * textScale * ProToken.maxLength;
}