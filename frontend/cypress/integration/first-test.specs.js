describe("First test", () => {
    it("should visit home page", () => {
      cy.visit("http://localhost:4200/");

      
    });
  });

  describe("Service visit", () => {
    it("should visit services page", () => {
      cy.visit("http://localhost:4200/service/1");
      });
  });

  describe("Book appointment", () => {
    it("should should show that booking is successful", () => {
        cy.get('[style="background-color: black; color: beige; border-radius: 6px;"]').click();
      });
  });
  ;