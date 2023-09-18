# Customer Portal - Configurable SSO Based Login

## Description 💡

This code sample demonstrates a robust JWT-based Single Sign-On (SSO) login flow for your customers accessing the customer portal. It effectively addresses crucial security concerns, including token impersonation and replay attacks. By following this example, you can seamlessly integrate the JWT-based SSO mechanism, reinforcing security while maintaining a smooth user experience. Safeguard user data and ensure secure access to the DevRev customer portal.

## Concepts 📚

Single Sign-On (SSO) is an authentication and authorization mechanism that allows users to access multiple applications and services with a single set of credentials (such as username and password) without the need to enter their credentials separately for each application. In other words, once a user logs in to one application, they gain access to all other connected applications without having to re-enter their login credentials.

By implementing SSO on DevRev, organizations can simplify user management,
reduce password-related help desk calls, and enforce consistent security
policies across their applications on DevRev.

## Authentication Flow ⚙️

1. **Activate Single Sign-On (SSO)**: Get in touch with DevRev Support to enable Single Sign-On for your customer portal. You'll need to provide a valid URL where your customers will be directed when they log in to the customer portal.

2. **Authenticate with Your Identity Provider (IdP)**: On the redirected page, use your Identity Provider (IdP) to verify the identity of the customer and retrieve their user information.

3. **Issue a one-time Token**: Share the user traits you've gathered with
   DevRev's Auth Service to obtain a one-time token for the session. You'll need your
   organization's Application Access Token (AAT) for this step.

4. **Enable Auto Membership**: With this token exchange, auto membership is enabled. If a user with the provided traits is not already in our system, we will automatically create a new user with the specified traits and generate the one-time token for them.

5. **Redirect to Customer Portal**: To complete the process, redirect the user to
   the customer portal, including the one-time token in the query parameters. Use
   this format for the URL:
   `http://support.devrev.ai/<your-org-slug>/callback/sso?jwt=<onetime-token>`

## The Onetime Token 🛡️

In our commitment to security and user protection, we use the one-time token to address the potential security risks associated with JWT-based authentications openly exposed in URLs. Some of the critical security concerns we aim to mitigate include:

1. [Token Impersonation](https://medium.com/@talsec/protecting-your-api-from-app-impersonation-token-hijacking-guide-and-mitigation-of-jwt-theft-48e744b76327)
2. [Replay
   Attacks](https://auth0.com/docs/secure/security-guidance/prevent-threats)

To bolster security, the one-time token possesses the following key features:

1. **Short-Lived Session**: The token has a brief lifespan of 5 minutes and should be used for redirection to the customer portal within this window. If not utilized within this time frame, a new token must be requested.

2. **Guaranteed One-Time Exchange**: Upon redirection to the customer portal, the one-time token can be exchanged for an authenticated session on the customer portal only once. Any attempt to reuse the token will result in a denied request, and any previously authenticated sessions associated with that token may be revoked to maintain a high level of security.

With these guidelines, you can ensure the security of your customer interactions with the DevRev platform, protecting their data and maintaining a trusted environment.
