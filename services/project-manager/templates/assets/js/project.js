// Fetch and display projects
function fetchProjects() {
    $.ajax({
      url: "http://localhost:8080/projects",
      type: "GET",
      success: function(data) {
        displayProjects(data);
      },
      error: function(error) {
        console.error("Error fetching projects:", error);
      }
    });
  }
  
  function displayProjects(projects) {
    var projectsDiv = document.getElementById("projects");
    projectsDiv.innerHTML = "";
  
    projects.forEach(function(project) {
      var projectDiv = document.createElement("div");
      projectDiv.innerHTML = "<h2>" + project.name + "</h2>";
  
      var createTaskForm = document.createElement("form");
      createTaskForm.innerHTML = '<input type="hidden" name="projectId" value="' + project.id + '">' +
        '<input type="text" name="name" placeholder="Task Name" class="border border-gray-400 rounded-l px-4 py-2">' +
        '<button type="submit" class="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded-r">Create Task</button>';
  
      createTaskForm.addEventListener("submit", function(e) {
        e.preventDefault();
        createTask(project.id);
      });
  
      projectDiv.appendChild(createTaskForm);
      projectsDiv.appendChild(projectDiv);
    });
  }
  
  function createProject() {
    var projectName = document.getElementById("projectName").value;
    $.ajax({
      url: "http://localhost:8080/createProject",
      type: "POST",
      data: { name: projectName },
      success: function() {
        fetchProjects();
        document.getElementById("projectName").value = "";
      },
      error: function(error) {
        console.error("Error creating project:", error);
      }
    });
  }
  
  function createTask(projectId) {
    var taskName = prompt("Enter Task Name:");
    if (!taskName) return;
  
    $.ajax({
      url: "http://localhost:8080/createTask",
      type: "POST",
      data: { name: taskName, projectId: projectId },
      success: function() {
        fetchProjects();
      },
      error: function(error) {
        console.error("Error creating task:", error);
      }
    });
  }
  
  // Fetch projects on page load
  fetchProjects();
  
  // Event listener for create project form submission
  document.getElementById("createProjectForm").addEventListener("submit", function(e) {
    e.preventDefault();
    createProject();
  });
  