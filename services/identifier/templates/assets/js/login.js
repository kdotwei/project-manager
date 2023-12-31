$(document).ready(function() {
    $('#loginForm').on('submit', function(e) {
        $('#loginError').hide();
        e.preventDefault();

        var userData = {
            username: $('#loginForm input[name="username"]').val(),
            password: $('#loginForm input[name="password"]').val()
        };

        $.ajax({
            type: "POST",
            url: "/identifier/login",
            contentType: "application/json",
            data: JSON.stringify(userData),
            success: function(response) {
                console.log(response);
                window.location.href = '/'; // Redirect to the protected page
            },
            error: function(xhr, status, error) {
                var response = JSON.parse(xhr.responseText);
                console.log(xhr.responseText);
                $('#loginError').text("Login failed: " + response.error).show();
            }
        });
    });
});
