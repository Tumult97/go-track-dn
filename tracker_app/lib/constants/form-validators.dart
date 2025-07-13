class FormValidators {
  static String? Function(dynamic) get notEmptyValidator => (value) {
    if (value == null || value.isEmpty) {
      return 'This field is required';
    }
    return null;
  };

  static String? Function(dynamic) get nonEmptyCurrencyValidator => (value) {
    var numberValue = double.tryParse(value);

    if (value == null || value.isEmpty || numberValue == null) {
      return 'This field is required';
    }

    if (numberValue < 0) {
      return 'Value cannot be negative';
    }

    return null;
  };
}