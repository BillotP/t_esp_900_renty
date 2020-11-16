import { createDecorator, VueDecorator } from 'vue-class-component';
import { PropOptions } from 'vue';

export type Constructor = new(...args: any[]) => any;

export function Props(options: (PropOptions | Constructor[] | Constructor) = {}): VueDecorator {
    return createDecorator((componentOptions, k) => {
        (componentOptions.props || (componentOptions.props = {}) as any)[k] = options;
    });
}
