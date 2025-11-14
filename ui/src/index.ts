import { LitElement, html, css } from 'lit';
import { customElement, property } from 'lit/decorators.js';

/**
 * A simple "Hello World" custom element.
 *
 * @element hello-world-element
 */
@customElement('hello-world-element')
export class HelloWorldElement extends LitElement {
  /**
   * The name to say hello to.
   * @attr
   */
  @property({ type: String })
  name = 'World';

  static styles = css`
    :host {
      display: block;
      border: solid 1px gray;
      padding: 16px;
      max-width: 800px;
    }
  `;

  render() {
    return html`<h1>Hello, ${this.name}!</h1>`;
  }
}
