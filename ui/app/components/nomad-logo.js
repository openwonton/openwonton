/**
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: MPL-2.0
 */

import Component from '@glimmer/component';
import config from 'nomad-ui/config/environment';

export default class NomadLogoComponent extends Component {
  get logoSrc() {
    const rootURL = config.rootURL || '/';
    const normalized = rootURL.endsWith('/') ? rootURL : rootURL + '/';
    return normalized + 'images/logo_openwonton.png';
  }
}
