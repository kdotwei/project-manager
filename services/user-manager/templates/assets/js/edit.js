// assets/js/edit.js

$(document).ready(function() {
    var userId = $('#userId').data('user-id');
    var getUserApiUrl = '/user-manager/api/users/' + userId;
    var putUserApiUrl = '/user-manager/api/users/' + userId + "/update";

    console.log(getUserApiUrl);
    console.log(putUserApiUrl);

    // Load existing user data
    $.ajax({
        url: getUserApiUrl,
        type: 'GET',
        dataType: 'json',
        success: function(user) {
            $('#username').val(user.user.username);
        },
        error: function(error) {
            // Error handling
            var errorMsg = error.responseJSON && error.responseJSON.error ? error.responseJSON.error : "Unknown error occurred, please try again.";
            $('#error-message').text(errorMsg).removeClass('hidden');
            console.error('Error get user:', error);
        }
    });

    // Handle form submission
    $('#editForm').submit(function(event) {
        event.preventDefault();
        var username = $('#username').val();
        var password = $('#password').val();

        var userData = {
            id: userId,
            username: username,
            password: password
        };

        $.ajax({
            type: 'PUT',
            url: putUserApiUrl,
            contentType: 'application/json',
            data: JSON.stringify(userData),
            success: function(response) {
                // Success handling
                console.log('User updated successfully:', response);
                window.location.href = '/user-manager/users';
            },
            error: function(error) {
                // Error handling
                var errorMsg = error.responseJSON && error.responseJSON.error ? error.responseJSON.error : "Unknown error occurred, please try again.";
                $('#error-message').text(errorMsg).removeClass('hidden');
                console.error('Error updating user:', error);
            }
        });
    });
});
