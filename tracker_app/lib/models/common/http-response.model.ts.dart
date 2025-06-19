class RequestResponse<T> {
  final int httpStatusCode;
  final T? data;
  final String? message;

  bool get isSuccess => httpStatusCode == 200 ? true : false;

  RequestResponse({required this.httpStatusCode, this.data, this.message});

  factory RequestResponse.success(T data) {
    return RequestResponse(httpStatusCode: 200, data: data);
  }

  factory RequestResponse.error(String message, int httpStatusCode) {
    return RequestResponse(httpStatusCode: httpStatusCode, message: message);
  }
}