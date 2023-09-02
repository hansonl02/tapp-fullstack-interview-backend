**Introduction**

Congratulations on making it to the coding interview for the software engineer role at StyleAI! Your task is to develop a new feature for our website editor: domain searching. Once a customer purchases a Style subscription, they will be able to choose a domain for hosting their new website. Your responsibility is to develop the user interface for the domain selection view and implement the backend function for querying available domains.

**Task Overview**

The task is split up into two parts: the frontend (https://github.com/tapp-ai/tapp-fullstack-interview-react) and the backend (https://github.com/tapp-ai/tapp-fullstack-interview-backend)

_Note: Refer to the design drawing below for the user flow_.

Tasks for the frontend:

- Develop the initial modal that appears when the user clicks "Find Domain!" to input their business name (`components/ExampleModal/index.jsx`).
- Design the modal that shows the available domains, which opens after the user inputs their business name (`components/ExampleModal/index.jsx`).
- Implement the necessary logic to handle the opening and closing of both modals in the Site (`components/Site/index.jsx`).
- When the user selects a domain, close the modal and display the chosen domain in the center of the Site (`components/Site/index.jsx`).

Tasks for the backend:

- Develop an endpoint (`getdomains.go`) that receives a business name as input, utilizes GPT to generate a list of potential domains, and then checks the availability of each domain. Finally, return a list of available domains to the frontend.
- Take note that the response from GPT is in string format, so you may need to parse it into a usable array.
- Utilize the following endpoint (http://domains.usestyle.ai/api/v1/availability?domain={DOMAIN}) to check the availability of a domain.

The files that you will be working on include:

Frontend:

- `components/Site/index.jsx`
- `components/ExampleModal/index.jsx` (referencing the sample Modal code)
- `components/…` (if you want to create any new components, ex. new modals)

Backend:

- `main.go`
- `getdomains.go`

**Time Expectation**

This task is designed to take an experienced developer roughly 1 - 2 hours. However, since you may not be familiar with the structure of this repository and our coding practices, you may need more time. It is totally okay if you take more time.

**Resource Usage**

You are free to use any standard internet resources like StackOverflow and Golang documentation, with the exception of language models like ChatGPT. If you use any external resource, please leave comments referencing where you got help from along the way.

**Tips for Getting Started**

Before getting started with the task, you should take the time to understand the repository's structure and coding practices. You can then proceed with the following steps:
_Note: Make sure you have npm (https://docs.npmjs.com/downloading-and-installing-node-js-and-npm) and Go (https://go.dev/doc/install) installed_

- First, get the backend up and running and connected to the frontend. To do this, run `go mod tidy` in the terminal to install all necessary go packages and `go run .` to start the backend locally.
- Then, in the frontend project root directory run `npm install` to install all react packages and `npm start` to start the frontend locally.
- To test the connection between the frontend and the backend is working, click on the "Check backend connection" button that should render on the frontend. You should receive a pop-up that says: "Success! The backend is properly connected." If this doesn't work, one common cause is the localhost port of the backend is incorrectly configured in the front end, so replace with the proper local url (`constants.js`)
- Feel free to try the sample "summarize" backend function using curl, Postman, Insomnia, or another API platform. This queries the OpenAI API to summarize some text.

If you encounter any issues or get stuck along the way, please reach out to your interviewer for assistance.

**Conclusion**

This task is designed to test your ability to work with React components, backend logic, and connecting full stack applications together. We wish you the best of luck in completing this task and look forward to reviewing your work.

**User Flow**

![User Flow](https://github.com/tapp-ai/tapp-fullstack-interview-backend/assets/56799300/d660f8d8-5d6d-4996-8a70-f1552b5ecbac)
