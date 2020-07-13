# Tort

Tort is a simple assertions library to help write Go unit tests.  It's early on, with only a 
smattering of assertions implemented, and no test cases.

## Example

Here's a sample of standard test code:

    func TestCreateTask(t *testing.T) {
        var details = "This is a sample task"
        var user = "jdoe@nowhere.com"
    
        task, err := domain.CreateTask(details, user)
        if err != nil {
            t.Fatalf("Unable to create task %s", details)
        }

        if task.ID == "" {
            t.Error("Expected task to have an identifier")
        }
        
        if task.Details != details {
            t.Errorf("Expected details to be %s; was %s", details, task.Details)
        }

        created := task.GetStatus(domain.StatusCreated)

        if s.Type != domain.StatusCreated {
        	t.Errorf("Expected status to be %s; was %s", domain.StatusCreated, s.Type)
        }

        if s.Updated.IsZero() || s.Updated.Since() > time.Minute {
            t.Error("Created timestamp is missing or incorrect")
        }

        if s.UpdatedBy != user.ID {
        	t.Errorf("Expected creator to be %s; was %s", user.ID, s.UpdatedBy)
        }
    }

Here's the same test sample, but with Tort:

    func TestCreateTask(t *testing.T) {
        assert := tort.For(t)

        var details = "This is a sample task"
        var user = "jdoe@nowhere.com"
    
        task, err := domain.CreateTask(details, user)
        assert.When("creating a task").Error(err).IsNil()
        assert.Struct(task).String("ID").NotBlank()
        assert.Struct(task).String("Details").Equals(details)

        created := task.GetStatus(domain.StatusCreated)
        assert.Struct(status).String("Type").Equals(domain.StatusCreated)
        assert.Struct(status).Time("Updated").IsSet()
        assert.Struct(status).Time("Updated").Within(time.Minute)
        assert.Struct(status).String("UpdatedBy").Equals(user.ID)
    }

Not only does it reduce the line count, but Tort adds some behaviors to allow contract-like behavior
on structs.  In the above, `assert.Struct(task).String("Details")` checks for the `Details` property
to exist on the `domain.Task` struct, and will fail abruptly if it's not there or not a string. 

## Assertions

The following outlines the assertions available for each data type.

### Errors

* `Nil`
* `NotNil`
* `Equals`
* `NotEquals`

### Booleans

* `IsTrue`
* `IsFalse`

### Integers and Floats

All integer types are supported:  `int`, `int8`, `int16`, `int32`, and `int64`, as well as unsigned
types, e.g. `uint8`, `uint16`, etc.  The `float32` and `float64` types are supported with similar
functions.

* `Equals`
* `NotEquals`
* `GreaterThan`
* `LessThan`

### Strings

* `Blank` - the string contains nothing but whitespace
* `NotBlank` - the string contains characters other than whitespace
* `Empty` - the string is empty of all characters, including whitespace
* `NotEmpty` - the string contains one or more characters of any kind, including whitespace
* `Equals`
* `NotEquals`
* `Folds` - like `Equals`, but ignores case
* `NotFolds` - like `NotEquals`, but ignores case
* `Contains` - the string contains the value
* `NotContains` - the string doesn't contain the value
* `Matches` - the string matches the regular expression
* `NotMatches` - the string doesn't match the regular expression

### Time

May be used with values of the type `time.Time`, or any struct that fulfills the `tort.TimeAssertion`
interface and implements `func Assert() time.Time`.

* `IsSet` - the time value is set, i.e. is `time.IsZero()` not true?
* `IsNotSet`  - the time value is not set
* `Within` - is the time within a certain duration from now; useful for checking if the time was set recently
* `Before` - is the time before the given time
* `After` - is the time after the given time

### Duration

* `Equals`
* `NotEquals`
* `GreaterThan`
* `LessThan`

### Slices

* `Empty` - there are no elements in the slice
* `Length` - does the slice length equal the given value?
* `MoreThan` - there are more than this number of elements in the slice
* `FewerThan` - there are fewer than his number of elements in the slice

#### Assertions on the slice elements

Slices also support assertions on slice elements.  Identify the element type 
by index, then apply any assertions to it for that type.

    func TestSlices(t *testing.T) {
	    assert := tort.For(t)
    
	    slice := []int{3, 1, 4, 1, 5, 9, 2}
	    
	    assert.Slice(slice).Length(7)
	    assert.Slice(slice).MoreThan(6)
	    assert.Slice(slice).FewerThan(8)
	    
	    assert.Slice(slice).Int(0).Equals(3)
	    assert.Slice(slice).Int(4).GreaterThan(4)
    }
 
Supports the following types:

* `Bool`
* `Int`, `Int8`, `Int16`, `Int32`, `Int64`
* `Uint`, `Uint8`, `Uint16`, `Uint32`, `Uint64`
* `Float32`, `Float64`
* `String`
* `Time`
* `Duration`
* `Struct`

### Structs

Structs focus on testing the properties of the structs.  It checks whether or not the property 
exists on the struct, whether or not it's the correct type, and then do the assertions match.

#### Assertions available on the struct itself

* `Nil` - is the struct a pointer, and is it nil?
* `NotNil` - is the pointer nil?  

#### Assertions on the properties

These assertions assert the named property is of the correct type.  If not, generates a fatal error.
Otherwise returns an object for testing the assertions based on the property type.

For example:

    type User struct {
        Email    string
        Password string
        Age      uint
    }
    
    func TestUser(t *testing.T) {
        assert := tort.For(t)
        
        u, err := NewUser("jdoe@nowhere.com", "mypassword", 25)
        assert.When("creating a new user").Error(err).IsNil()
        
        assert.Struct(u).String("Email").Equals("jdoe@nowhere.com")
        assert.Struct(u).Uint("Age").Equals(25)
        
        // This should have been encrypted!
        assert.Struct(u).String("Password").NotEquals("mypassword")
    }
    
* `Bool`
* `Int`, `Int8`, `Int16`, `Int32`, `Int64`
* `Uint`, `Uint8`, `Uint16`, `Uint32`, `Uint64`
* `Float32`, `Float64`
* `String`
* `Time`
* `Duration`
* `Struct`

Note with `Struct`, you can drill down into structs.  If a property doesn't exist, a fatal error
occurs.

    assert.Struct(user).Struct("Address").String("City").Equals("New York")

## The `When` Function

The `When` function provides some context around a test.  It may be placed on the assertions at the
start, for example:

    assert.When("updating the document").Error(err).IsNil()
    assert.When("retrieving a user").Struct(user).String("Email").Matches(`\w+@\w+\.\w+`)
    
You may also save the result of the `When` call and reuse it repeatedly:

    assert = assert.When("fetching a record")
    assert.Struct(record).String("name").Equals("Fun in the Sun") // will log "fetching a record" on error
    
The string value passed to `When` is included in the error output.
