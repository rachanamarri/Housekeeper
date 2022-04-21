

describe("First Test", () => {

    it("must be able to visit the home page", () => {

      cy.visit("http://localhost:4200/");

    });
    
  });

  describe("Service Visit", () => {

    it("must be able to visit the services page", () => {

      cy.visit("http://localhost:4200/service/1");

      });

  });

  describe("Book Appointment", () => {

    it("must be able to show that booking is successful", () => {

        cy.get('[style="background-color: black; color: beige; border-radius: 6px;"]').click();

      });

  });

  ;

