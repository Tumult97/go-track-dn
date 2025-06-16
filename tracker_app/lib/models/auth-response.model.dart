class AuthResponse {
  final int id;
  final String firstName;
  final String lastName;
  final String email;
  final DateTime createdAt;
  final String? accessToken;
  final String? refreshToken;
  final String message;

  AuthResponse({
    required this.id,
    required this.firstName,
    required this.lastName,
    required this.email,
    required this.createdAt,
    this.accessToken,
    this.refreshToken,
    required this.message,
  });

  factory AuthResponse.fromJson(Map<String, dynamic> json) {
    return AuthResponse(
      id: json['id'] as int,
      firstName: json['first_name'] as String,
      lastName: json['last_name'] as String,
      email: json['email'] as String,
      createdAt: DateTime.parse(json['created_at'] as String),
      accessToken: json['access_token'] as String?,
      refreshToken: json['refresh_token'] as String?,
      message: json['message'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'first_name': firstName,
      'last_name': lastName,
      'email': email,
      'created_at': createdAt.toIso8601String(),
      if (accessToken != null) 'access_token': accessToken,
      if (refreshToken != null) 'refresh_token': refreshToken,
      'message': message,
    };
  }
}
