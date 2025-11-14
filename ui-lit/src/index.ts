import {LitElement, html} from 'lit';
import {customElement, property} from 'lit/decorators.js';

@customElement('core-element-template-lit')
export class CoreElementTemplateLit extends LitElement {
  @property()
  title = 'core-element-template';

  render() {
    return html`<h1>Hello, ${this.title}</h1>`;
  }
}
