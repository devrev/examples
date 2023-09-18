const AAT = "your-private-AAT";

// Handles the click event for the Login button.
const clickButton = async () => {
  // Fetch a onetime session token with the user identifiers provided by your IDP.
  const res = await fetch("https://api.devrev.ai/auth-tokens.create", {
    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
      Authorization: AAT,
      Host: "localhost:4100",
    },
    body: JSON.stringify({
      grant_type: "urn:ietf:params:oauth:grant-type:token-exchange",
      requested_token_type:
        "urn:devrev:params:oauth:token-type:session:onetime",
      subject_token_type: "urn:devrev:params:oauth:token-type:revinfo",
      rev_info: {
        user_ref: "<user-ref>", // Unique identifier of the user on the IDP.
        org_ref: "shipsy.com", // External reference of the rev org of the user.
        user_traits: {
          email: "john.smith@shipsy.com", // Email ID of the user.
        },
      },
    }),
  });
  const content = await res.json();
  console.log(content);

  // Redirect the user to the support portal with the onetime token in the query params.
  window.location.href = `http://support.devrev.ai/<your-org-slug>/callback/sso?jwt=${content?.access_token}`;
};
