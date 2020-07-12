# Tort

Tort is a simple assertions library to help write Go unit tests.  It is very early on, and Tort
lacks test cases of its own.

Here's a sample of standard test code:

    func TestCreateTask(t *testing.T) {
        var details = "This is a sample task"
    
        task, err := domain.CreateTask(details, user.ID)
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
    
        task, err := domain.CreateTask(details, user.ID)
        assert.When("creating a task").Error(err).IsNil()
        assert.Struct(task).String("ID").NotBlank()
        assert.Struct(task).String("Details").Equals(details)

        created := task.GetStatus(domain.StatusCreated)
        assert.Struct(status).String("Type").Equals(domain.StatusCreated)
        assert.Struct(status).Time("Updated").Set()
        assert.Struct(status).Time("Updated").Within(time.Minute)
        assert.Struct(status).String("UpdatedBy").Equals(user.ID)
    }

Not only does it reduce the line count, but Tort adds some behaviors to allow contract-like behavior
on structs.  In the above, `assert.Struct(task).String("Details")` checks for the `Details` property
to exist on the `domain.Task` struct, and will fail abruptly if it's not there or not a string. 
