function makeGetRequest(url, callback) {
    const xhr = new XMLHttpRequest();
    xhr.open('GET', url, true);

    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
                callback(xhr.responseText);
            } else {
                console.error('Request error:', xhr.status, xhr.statusText);
            }
        }
    };

    xhr.send();
}

function loadTestPreviews() {
    const apiUrl = 'http://localhost:8080';

    makeGetRequest(apiUrl, function (response) {
        const tests = JSON.parse(response);
        const testPreviewsContainer = document.getElementById('test-previews');

        tests.forEach(function (test) {
            const testPreview = document.createElement('div');
            testPreview.classList.add('test-preview');
            const id = 'test_id_' + test.id;
            testPreview.setAttribute('id', id);

            testPreview.addEventListener('click', function () {
                console.log('Clicked ' + id);
                getTest(test.id);
            });

            testPreview.innerHTML = `
                <h3>${test.title}</h3>
                <p>${test.description}</p>
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

    makeGetRequest(apiUrl, function (response) {
        const test = JSON.parse(response);
        console.log(test)

        const question_count = test.questions.length;
        document.getElementById("test-title").textContent = test.preview.title;
        document.getElementById("test-description").textContent = test.preview.description;
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
                console.log(answer_ids);
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
    questionText.textContent = questionData.question_text;

    const optionsList = document.getElementById("options-list");
    optionsList.innerHTML = ""; 

    const optionsForm = document.getElementById("options-form");
    questionData.options.forEach((option, index) => {
        const li = document.createElement("li");
        const input = document.createElement("input");
        const label = document.createElement("label");

        input.type = "radio";
        input.name = "option"; 
        input.value = option.option_id; 
        input.id = `option-${index}`;

        label.textContent = `${index + 1}. ${option.option_text}`;
        label.htmlFor = `option-${index}`;


        li.appendChild(input);
        li.appendChild(label);
        optionsList.appendChild(li);
    });
}


window.addEventListener('load', loadTestPreviews);

