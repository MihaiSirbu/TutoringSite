


document.addEventListener('DOMContentLoaded', function() {
    // Replace these values with actual data from your database
    const lessonData = {
        title: 'Lesson 5: Finding X',
        date: 'January 20, 2020',
        description: 'During this lesson, we reviewed how to solve for x using basic equations.',
        exercises: [
            'x + 5 = 10',
            // Add more exercises as needed
        ]
    };

    // Populate lesson data
    document.getElementById('lesson-title').textContent = lessonData.title;
    document.getElementById('lesson-date').textContent = lessonData.date;
    document.getElementById('lesson-description').textContent = lessonData.description;

    // Load exercises
    const exercisesContainer = document.getElementById('exercises');
    lessonData.exercises.forEach(function(exercise, index) {
        const exerciseElement = document.createElement('div');
        exerciseElement.className = 'exercise';
        exerciseElement.innerHTML = `
            <label for="exercise-${index}">${exercise}</label>
            <input type="text" id="exercise-${index}" name="exercise-${index}" placeholder="Your answer">
        `;
        exercisesContainer.appendChild(exerciseElement);
    });

    // Handle form submission
    document.getElementById('feedback-form').addEventListener('submit', function(event) {
        event.preventDefault();
        // Here you would handle the submission, e.g., send data to the server
        alert('Feedback submitted!');
    });
});
