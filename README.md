<h1>zaf-telco</h1>

<h3>About</h3>
[Include a brief description or introduction of the project here. Explain what the project does, its purpose, and any key features or functionalities.]
This project is my take on implementing a CQRS/Event Spurcing in a Microservice architecture. There will not be any real business logic here but instead, I will focus more on the technicalities of how to implement a service using the CQRS/Event Source pattern.

In this project, I will simulate a working example of an MVNO (Mobile Virtual Network Operator) company's backend architecture which would include multiple domain specific services. This is mainly to gain some knowledge on working with a CQRS and Event Sourcing architechture pattern. These services would communicate via GRPC or through an Event driven system. 

These services include:
 <ol>
  <li>Telco (Current Service) Deals with a user's telco accounts</li>
  <li>User Service (Deals with creating a user's KYC profile in the company)</li>
  <li>Billing (Deals with anything realated to billing. From usage details to add-on purchases)</li>
  <li>Inventory Management Service (Deals with managing all data/voice service add-ons)</li>
  <li>Events Service (Deals with managing all events that happens within the system. Captures all events, creates an event, sends the event to responsible services and stores the event in an aggregate for the mobile/frontend to present to the users)</li>
  <li>BFF or backend for the frontend (Deals with fetching the presentation for the mobile/frontend app for the users)</li>
</ol>

<h3>Installation</h3>
[Provide instructions on how to install and set up the project. Include any dependencies or prerequisites needed, along with the necessary commands or steps.]

<h3>Usage</h3>
[Explain how to use the project and provide examples or code snippets if applicable. Include any important details or considerations that users should be aware of.]

<h3>Features</h3>
[List the main features or functionalities of the project.]

Contributing
[Explain how others can contribute to the project, including guidelines for submitting bug reports, feature requests, or pull requests. Include any specific instructions or coding standards.]

License
[Specify the license under which the project is distributed. Include any terms or conditions associated with the license.]

Credits
[List any individuals, organizations, or resources that deserve credit for contributing to the project.]

Contact
[Provide contact information for the project maintainer or team. Include email addresses, social media handles, or any other relevant details.]

Acknowledgements
[Optional: Acknowledge any external resources, libraries, or code snippets that were used in the project.]

Disclaimer
[Include any necessary disclaimers or legal notices related to the project.]

Changelog
[Optional: Provide a summary of changes for each version of the project, highlighting significant updates, bug fixes, or new features.]