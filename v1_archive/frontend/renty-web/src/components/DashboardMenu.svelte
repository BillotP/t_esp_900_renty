<script lang="ts">
  export let segment: string;
</script>

<aside class="menu is-primary" style="padding-top: 10px;">
  <p class="menu-label">General</p>
  <ul class="menu-list">
    <li>
      <a class:is-active={segment === undefined} href="dashboard">
        Dashboard
      </a>
    </li>
    <li>
      <a class:is-active={segment === 'rentoffers'} href="dashboard/rentoffers">
        RentOffers
      </a>
    </li>
  </ul>
  <p class="menu-label">Administration</p>
  <ul class="menu-list">
    <li>
      <a class:is-active={segment === 'documents'} href="dashboard/documents">
        My documents
      </a>
    </li>
    <li>
      <a class:is-active={segment === 'selection'} href="dashboard/selection">
        My selection
      </a>
      <ul>
        <li>
          <a
            class:is-active={segment === 'offers'}
            href="dashboard/selection/offers">
            Offers
          </a>
        </li>
        <li>
          <a
            class:is-active={segment === 'progress'}
            href="dashboard/selection/progress">
            Progress
          </a>
        </li>
        <li>
          <a
            class:is-active={segment === 'appointment'}
            href="dashboard/selection/appointment">
            Appointments
          </a>
        </li>
      </ul>
    </li>
  </ul>
  <p class="menu-label">Transactions</p>
  <ul class="menu-list">
    <li>
      <a
        class:is-active={segment === 'payments'}
        href="dashboard/transactions/payments">
        Payments
      </a>
    </li>
    <li>
      <a
        class:is-active={segment === 'payments'}
        href="dashboard/transactions/payments">
        Transfers
      </a>
    </li>
    <li>
      <a
        class:is-active={segment === 'balance'}
        href="dashboard/transactions/balance">
        Balance
      </a>
    </li>
  </ul>
</aside>
