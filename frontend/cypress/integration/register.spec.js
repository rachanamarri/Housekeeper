describe("Register test", () => {
    it("should register a user", () => {
      cy.visit("http://localhost:4200/register");

      cy.get("#registerUName").type("Test User");
      cy.get("#registerEmail").type("TestUser@test.com");
      cy.get("#registerPswd").type("password");
      cy.get("#registerAddress").type("adress is this");
      cy.get("#registerSubmit").click();
      
    });
  });