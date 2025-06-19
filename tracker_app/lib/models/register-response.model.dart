class RegisterResponse {
  final int id;
  final String firstName;
  final String lastName;
  final String email;
  final DateTime createdAt;
  final String? accessToken;
  final String? refreshToken;
  final String message;

  RegisterResponse({
    required this.id,
    required this.firstName,
    required this.lastName,
    required this.email,
    required this.createdAt,
    this.accessToken,
    this.refreshToken,
    required this.message,
  });

  factory RegisterResponse.fromJson(Map<String, dynamic> json) {
    return RegisterResponse(
      id: json['id'],
      firstName: json['first_name'],
      lastName: json['last_name'],
      email: json['email'],
      createdAt: DateTime.parse(json['created_at']),
      accessToken: json['access_token'],
      refreshToken: json['refresh_token'],
      message: json['message'],
    );
  }
}