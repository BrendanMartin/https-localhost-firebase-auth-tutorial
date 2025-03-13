This is the completed example for the tutorial [Setting up Firebase Redirect Auth on localhost with HTTPS](https://brmartin.com/firebase-redirect-auth-https-localhost/).

In the project root, generate an SSL certificate with the following command:
```bash
openssl req -x509 -newkey rsa:4096 -keyout localhost.key -out localhost.crt -sha256 -days 36500 -nodes -subj "/C=US/CN=localhost" -addext "subjectAltName = DNS:localhost"
```

Add `localhost.crt` to your OS's trusted certificates.

Install javascript dependencies with `npm install`.

Create a `firebase-config.js` with your project's config values.

Run the go server with `air`.

Navigate to `https://localhost:8082` to see a fully functioning Google sign-in button.
