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
});
