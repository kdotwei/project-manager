// assets/js/user.js

$(document).ready(function() {
    var userId = $('#userId').data('user-id');
    var userApiUrl = '/user-manager/api/users/' + userId;
    console.log(userApiUrl);

    $.ajax({
        url: userApiUrl,
        type: 'GET',
        dataType: 'json',
        success: function(user) {
            user = user.user;
            $('#username').text(user.username);
            $('#create_time').text(new Date(user.CreatedAt).toLocaleDateString() + " " + new Date(user.CreatedAt).toLocaleTimeString());
            $('a[href="#edit"]').attr('href', '/user-manager/users/' + userId + '/edit');
            $('a[href="#delete"]').attr('href', '/user-manager/api/users/' + userId + '/delete');
        },
        error: function(error) {
            var errorMsg = error.responseJSON && error.responseJSON.error ? error.responseJSON.error : "Unknown error occurred, please try again.";
            $('#error-message').text(errorMsg).removeClass('hidden');
            console.error('Error fetching user:', error);
        }
    });

    // Delete user event
    $('#deleteButton').click(function() {
        if (confirm('Are you sure you want to delete this user?')) {
            $.ajax({
                url: userApiUrl + '/delete',
                type: 'DELETE',
                success: function(response) {
                    alert('User deleted successfully');
                    window.location.href = '/user-manager/users'; // Redirect to users list
                },
                error: function(error) {
                    var errorMsg = error.responseJSON && error.responseJSON.error ? error.responseJSON.error : "Error occurred during deletion.";
                    $('#error-message').text(errorMsg).removeClass('hidden');
                    console.error('Error deleting user:', error);
                }
            });
        }
    });
});
