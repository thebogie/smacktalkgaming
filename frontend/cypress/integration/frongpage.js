describe('smacktalk front page', () => {
	beforeEach(() => {
		cy.visit('http://192.168.86.45:3000')
	});

	it('has the correct title', () => {
		cy.contains('<title>', 'Smack Talk Gaming')
	});

	it('navigates to Log in', () => {
		cy.get('nav a').contains('Log in').click();
		cy.url().should('include', '/login');
	});

});