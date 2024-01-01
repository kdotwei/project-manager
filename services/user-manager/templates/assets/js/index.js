// assets/js/index.js

$(document).ready(function() {
    $.ajax({
        url: '/user-manager/api/users',
        type: 'GET',
        dataType: 'json',
        success: function(response) {
            var tbodyContent = '';
            response.users.forEach(function(user) {
                tbodyContent += `
                    <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600">
                        <td class="px-6 py-4 text-gray-900 whitespace-nowrap dark:text-white">
                            ${user.ID}
                        </td>
                        <td class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                            ${user.username}
                        </td>
                        <td class="px-6 py-4">
                            ${new Date(user.CreatedAt).toLocaleDateString()} ${new Date(user.CreatedAt).toLocaleTimeString()}
                        </td>
                        <td class="px-6 py-4 text-right">
                            <button id="edit-${user.ID}" class="font-medium text-blue-600 dark:text-blue-500 hover:underline mr-3">Edit</button>
                            <button id="delete-${user.ID}" class="font-medium text-red-600 dark:text-red-500 hover:underline">Delete</button>
                        </td>
                    </tr>
                `;
            });
            // Use the ID to target the table body and update its content
            $('#usersTableBody').html(tbodyContent);
        },
        error: function(error) {
            var response = JSON.parse(error.responseText);
            var errorAlert = $('#error-message');
            errorAlert.text(response.error);
            errorAlert.removeClass('hidden');
            console.error('Error fetching users:', error);
        }
    });
});
