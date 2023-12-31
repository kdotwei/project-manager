$(document).ready(function() {
    $('#loginForm').on('submit', function(e) {
        $('#loginError').hide();
        e.preventDefault();

        var userData = {
            username: $('#loginForm input[name="username"]').val(),
            password: $('#loginForm input[name="password"]').val()
        };

        var dataToSend = JSON.stringify(userData);
        console.log("Data to send:", dataToSend);

        $.ajax({
            type: "POST",
            url: "/identifier/login",
            contentType: "application/json",
            data: dataToSend,
            success: function(response) {
                console.log("Success!");
                console.log(response);
                window.location.href = '/'; // Redirect to the protected page
            },
            error: function(xhr, status, error) {
                var response = JSON.parse(xhr.responseText);
                console.log("Error!");
                console.log(xhr.responseText);
                $('#loginError').text("Login failed: " + response.error).show();
            }
        });
    });
});
