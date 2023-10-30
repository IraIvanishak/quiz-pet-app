function makeRequest(url, method, data, callback) {
    const options = {
        method: method,
        headers: {
            'Content-Type': 'application/json'
        },
        credentials: 'include',  // Important for handling cookies
    };

    if (data) {
        options.body = JSON.stringify(data);
    }

    fetch(url, options)
        .then(response => {
            if (!response.ok) {
                return Promise.reject(`Request error: ${response.status} ${response.statusText}`);
            }
            return response.json();
        })
        .then(data => callback(data))
        .catch(error => console.error('Error:', error));
}

function sendAnswers(id, answer_ids) {
    const apiUrl = `http://localhost:8080/test?id=${id}`;
    makeRequest(apiUrl, 'POST', answer_ids, function (response) {
        console.log('Response data:', response);
        showResult(response);  // Assuming the response contains the result
    });
}

function showResult(result) {
    const resultSection = document.getElementById('result');
    resultSection.style.display = 'block';

    const testSection = document.getElementById('test');
    testSection.style.display = 'none';

    const questionSection = document.getElementById('question');
    questionSection.style.display = 'none';

    // Assuming result is a string like "3/4"
    resultSection.innerHTML = `<h2>Your Score: ${result}</h2>`;
}

function loadTestPreviews() {
    const apiUrl = 'http://localhost:8080';
    makeRequest(apiUrl, 'GET', null, function (response) {

        console.log(response);
        const tests = response;
        const testPreviewsContainer = document.getElementById('test-previews');

        tests.forEach(function (test) {
            const testPreview = document.createElement('div');
            testPreview.classList.add('test-preview');
            const id = 'test_id_' + test.test.ID;
            testPreview.setAttribute('id', id);

            testPreview.addEventListener('click', function () {
                console.log('Clicked ' + id);
                getTest(test.test.ID);
            });

            testPreview.innerHTML = `
            <h3>${test.test.Title}</h3>
            <p>${test.test.Description}</p>
            ${test.score !== -1 ? `<div><label>Score: </label><span class="score">${test.score}</span></div>` : ''}
            `;

            testPreviewsContainer.appendChild(testPreview);
        });
    });
}

function getTest(id) {
    document.getElementById("test-preview").style.display = "none";
    const testSection = document.getElementById("test");
    testSection.style.display = "block";
    const apiUrl = `http://localhost:8080/test?id=${id}`;

    makeRequest(apiUrl, 'GET', null, function (response) {
        const test = response;
        console.log(test)

        const question_count = test.questions.length;
        document.getElementById("test-title").textContent = test.preview.Title;
        document.getElementById("test-description").textContent = test.preview.Description;
        document.getElementById("questions-amount").textContent = question_count;
        const nextQuestionButton = document.getElementById("next-question-button");
        let currQuestion = 0;
        let answer_ids = [];
        nextQuestionButton.addEventListener("click", () => {
            const selectedOption = document.querySelector('input[name="option"]:checked');
            if (selectedOption) {
                const selectedAnswerId = selectedOption.value;
                answer_ids.push(selectedAnswerId)
            } else {
                console.log("No option selected");
                answer_ids.push('-1')
            }
            if (currQuestion == question_count - 2) {
                nextQuestionButton.textContent = "Submit";
            } else if (currQuestion == question_count - 1) {
                sendAnswers(id, answer_ids)
                return;
            }
            currQuestion++;
            renderQuestion(test.questions[currQuestion]);
        });

        const startButton = document.getElementById("start-button");
        startButton.addEventListener("click", function () {
            testSection.style.display = "none";
            renderQuestion(test.questions[0])
        });
    });
}

function renderQuestion(questionData) {
    const questionSection = document.getElementById("question");
    questionSection.style.display = "block";

    const questionText = document.getElementById("question-text");
    questionText.textContent = questionData.QuestionText.String;

    const optionsList = document.getElementById("options-list");
    optionsList.innerHTML = "";

    questionData.Options.RawMessage.forEach((option, index) => {
        const li = document.createElement("li");
        const input = document.createElement("input");
        const label = document.createElement("label");

        input.type = "radio";
        input.name = "option";
        input.value = index;  // Assuming option index is used as value
        input.id = `option-${index}`;

        label.textContent = `${index + 1}. ${option.option_text}`;
        label.htmlFor = `option-${index}`;

        li.appendChild(input);
        li.appendChild(label);
        optionsList.appendChild(li);
    });
}

window.addEventListener('load', loadTestPreviews);
