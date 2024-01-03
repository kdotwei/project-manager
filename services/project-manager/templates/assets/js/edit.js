// assets/js/edit.js

$(document).ready(function() {
    var projectId = $('#projectId').data('projectId');
    var getProjectApiUrl = '/project-manager/api/projects/' + projectId;
    var putProjectApiUrl = '/project-manager/api/projects/' + projectId + "/update";

    console.log(getProjectApiUrl);
    console.log(putProjectApiUrl);

    // Load existing project data
    $.ajax({
        url: getProjectApiUrl,
        type: 'GET',
        dataType: 'json',
        success: function(project) {
            $('#projectName').val(project.name);
        },
        error: function(error) {
            // Error handling
            var errorMsg = error.responseJSON && error.responseJSON.error ? error.responseJSON.error : "Unknown error occurred, please try again.";
            $('#error-message').text(errorMsg).removeClass('hidden');
            console.error('Error get project:', error);
        }
    });

    // Handle form submission
    $('#editForm').submit(function(event) {
        event.preventDefault();
        var projectName = $('#projectName').val();

        var projectData = {
            id: projectId,
            name: projectName,
        };

        $.ajax({
            type: 'PUT',
            url: putProjectApiUrl,
            contentType: 'application/json',
            data: JSON.stringify(projectData),
            success: function(response) {
                // Success handling
                console.log('Project updated successfully:', response);
                window.location.href = '/project-manager/projects';
            },
            error: function(error) {
                // Error handling
                var errorMsg = error.responseJSON && error.responseJSON.error ? error.responseJSON.error : "Unknown error occurred, please try again.";
                $('#error-message').text(errorMsg).removeClass('hidden');
                console.error('Error updating project:', error);
            }
        });
    });
});
