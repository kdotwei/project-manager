// assets/js/project.js

$(document).ready(function() {
  var projectId = $('#projectId').data('projectId');
  
  $('#createTaskForm').submit(function(event) {
      event.preventDefault();
      createTask(projectId);
  });

  loadTasks(projectId);

  function loadTasks(projectId) {
      $.ajax({
          url: `/project-manager/api/projects/${projectId}/tasks`,
          type: 'GET',
          dataType: 'json',
          success: function(response) {
              console.log(response);
              var tbodyContent = '';
              response.tasks.forEach(function(task) {
                  var taskRow = `
                      <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                          <td class="px-6 py-4 text-gray-900 whitespace-nowrap dark:text-white">${task.ID}</td>
                          <td class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">${task.name}</td>
                          <td class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">${task.status}</td>
                          <td class="px-6 py-4">${new Date(task.CreatedAt).toLocaleDateString()} ${new Date(task.CreatedAt).toLocaleTimeString()}</td>
                          <td class="px-6 py-4">${new Date(task.UpdatedAt).toLocaleDateString()} ${new Date(task.UpdatedAt).toLocaleTimeString()}</td>
                          <td class="px-6 py-4 text-right">
                              <button id="edit-${task.ID}" class="edit-btn font-medium text-blue-600 dark:text-blue-500 hover:underline mr-3">Edit</button>
                              <button id="delete-${task.ID}" class="delete-btn font-medium text-red-600 dark:text-red-500 hover:underline">Delete</button>
                          </td>
                      </tr>
                  `;
                  tbodyContent += taskRow;
              });
              $('#tasksTableBody').html(tbodyContent);

              // Click event listener for Edit buttons
              $('.edit-btn').click(function(event) {
                  event.stopPropagation(); // Prevent triggering the project-row click event
                  var taskId = $(this).attr('id').split('-')[1];
                  window.location.href = `/project-manager/projects/${projectId}/tasks/${taskId}/edit`;
              });
              
              // Click event listener for Delete buttons
              $('.delete-btn').click(function(event) {
                  event.stopPropagation(); // Prevent triggering the project-row click event
                  var taskId = $(this).attr('id').split('-')[1];
                  deleteTask(taskId);
              });
          },
          error: function(error) {
              // Error handling
              var error_message = error.responseJSON && error.responseJSON.error ? error.responseJSON.error : "Unknown error occurred, please try again.";
              $('#error-message').text(error_message).removeClass('hidden');
              console.error('Error deleting project:', error);
          }
      });
  }

  // Function to create a task
  function createTask(projectId) {
    var taskName = $('#task-name').val();
    var data = {
        Name: taskName,
        Status: ""
    };

    console.log(data);

    $.ajax({
        type: 'POST',
        url: `/project-manager/api/projects/${projectId}/tasks/create`, // API endpoint for task creation
        contentType: 'application/json',
        data: JSON.stringify(data),
        success: function(response) {
            loadTasks(projectId); // Reload tasks after creation
        },
        error: function(error) {
            // Error handling
            var error_message = error.responseJSON && error.responseJSON.error ? error.responseJSON.error : "Unknown error occurred, please try again.";
            $('#error-message').text(error_message).removeClass('hidden');
            console.error('Error deleting project:', error);
        }
    });
  }

  // Function to delete a task
  function deleteTask(taskId) {
    if (confirm('Are you sure you want to delete this task?')) {
        $.ajax({
            url: `/project-manager/api/projects/${projectId}/tasks/${taskId}/delete`, // API endpoint for task deletion
            type: 'DELETE',
            success: function(response) {
                loadTasks(projectId); // Reload tasks after deletion
            },
            error: function(error) {
                // Error handling
                var error_message = error.responseJSON && error.responseJSON.error ? error.responseJSON.error : "Unknown error occurred, please try again.";
                $('#error-message').text(error_message).removeClass('hidden');
                console.error('Error deleting project:', error);
            }
        });
    }
  }


  function getProjectIdFromURL() {
      const urlParams = new URLSearchParams(window.location.search);
      return urlParams.get('id');
  }
});
