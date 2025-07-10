<script setup>
const props = defineProps({
  id: {
    type: String,
    required: true,
  },
})
</script>

<template>
  <div class="card-body pc-component">
    <div
      :id="id"
      class="modal fade"
      tabindex="-1"
      role="dialog"
      :aria-labelledby="`${id}Label`"
      aria-hidden="true"
    >
      <div class="modal-dialog" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" :id="`${id}Label`">
              <slot name="title"></slot>
            </h5>
            <button
              type="button"
              class="text-lg flex items-center justify-center rounded w-7 h-7 text-secondary-500 hover:bg-danger-500/10 hover:text-danger-500"
              :data-pc-modal-dismiss="`#${id}`"
              aria-label="Close"
            >
              <i class="ti ti-x"></i>
            </button>
          </div>
          <div class="modal-body">
            <slot></slot>
          </div>
          <div class="modal-footer">
            <slot name="footer">
              <button type="button" class="btn btn-secondary" :data-pc-modal-dismiss="`#${id}`">
                Close
              </button>
            </slot>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
