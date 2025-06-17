class RefreshTokenResponse {
  final String accessToken;
  final String refreshToken;
  final String message;

  RefreshTokenResponse({
    required this.accessToken,
    required this.refreshToken,
    required this.message,
  });

  factory RefreshTokenResponse.fromJson(Map<String, dynamic> json) {
    return RefreshTokenResponse(
      accessToken: json['access_token'] as String,
      refreshToken: json['refresh_token'] as String,
      message: json['message'] as String,
    );
  }
}