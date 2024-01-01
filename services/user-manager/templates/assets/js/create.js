$(document).ready(function() {
	$('#submitForm').submit(function(event) {
	event.preventDefault();

	// Show the loading spinner
	$('#loading-spinner').removeClass('hidden');
	$('#submit-button').addClass('hidden');

	var username = $('#username').val();
	var password = $('#password').val();

	var userData = {
		username: username,
		password: password
	};

	$.ajax({
		type: 'POST',
		url: '/user-manager/api/users/create',
		contentType: 'application/json',
		data: JSON.stringify(userData),
		success: function(response) {
			console.log('User created successfully:', response);
			window.location.href = '/user-manager/users';
		},
		error: function(error) {
			var errorMsg = error.responseJSON && error.responseJSON.error ? error.responseJSON.error : "Unknown error occurred, please try again.";
			$('#error-message').text(errorMsg).removeClass('hidden');
			console.error('Error fetching users:', error);
		},
		complete: function() {
			// Hide the loading spinner
			$('#loading-spinner').addClass('hidden');
			$('#submit-button').removeClass('hidden');
		}
	});
	});
});
