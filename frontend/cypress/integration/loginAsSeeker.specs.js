describe("login test", () => {
    it("should login a user", () => {
      cy.visit("http://localhost:4200/login");

      cy.get("#loginEmail").type("mithali@123.com");
      cy.get("#loginpswrd").type("abcd123");
      cy.get("#loginAsSeeker").click();
      
    });
  });
