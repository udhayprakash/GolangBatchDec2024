# Assertions


assert.Equal(t, expected, actual, msg): Checks if expected is equal to actual.

assert.NotEqual(t, expected, actual, msg): Checks if expected is not equal to actual.

assert.Nil(t, value, msg): Checks if value is nil.

assert.NotNil(t, value, msg): Checks if value is not nil.

assert.True(t, condition, msg): Checks if condition is true.

assert.False(t, condition, msg): Checks if condition is false.

assert.Error(t, err, msg): Checks if err is not nil (i.e., an error occurred).

assert.NoError(t, err, msg): Checks if err is nil (i.e., no error occurred).