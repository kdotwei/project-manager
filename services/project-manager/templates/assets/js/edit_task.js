// assets/js/edit_task.js

$(document).ready(function() {
    var projectId = $('#projectId').data('projectId');
    var taskId = $('#taskId').data('taskId');
    var getTaskApiUrl = '/project-manager/api/projects/' + projectId + '/tasks/' + taskId;
    var putTaskApiUrl = '/project-manager/api/projects/' + projectId + '/tasks/' + taskId + "/update";

    console.log(getTaskApiUrl);
    console.log(putTaskApiUrl);

    // Load existing task data
    $.ajax({
        url: getTaskApiUrl,
        type: 'GET',
        dataType: 'json',
        success: function(task) {
            console.log(task.name);
            $('#taskName').val(task.name);
            $('#taskStatus').val(task.status);
        },
        error: function(error) {
            // Error handling
            var error_message = error.responseJSON && error.responseJSON.error ? error.responseJSON.error : "Unknown error occurred, please try again.";
            $('#error-message').text(error_message).removeClass('hidden');
            console.error('Error deleting project:', error);
        }
    });

    // Handle form submission
    $('#editTaskForm').submit(function(event) {
        event.preventDefault();
        var taskName = $('#taskName').val();
        var taskStatus = $('#taskStatus').val();

        var taskData = {
            id: taskId,
            name: taskName,
            status: taskStatus
        };

        $.ajax({
            type: 'PUT',
            url: putTaskApiUrl,
            contentType: 'application/json',
            data: JSON.stringify(taskData),
            success: function(response) {
                // Success handling
                console.log('Project deleted:', response);
                window.location.href = `/project-manager/projects/${projectId}/tasks`;
            },
            error: function(error) {
                // Error handling
                var error_message = error.responseJSON && error.responseJSON.error ? error.responseJSON.error : "Unknown error occurred, please try again.";
                $('#error-message').text(error_message).removeClass('hidden');
                console.error('Error deleting project:', error);
            }
        });
    });
});
