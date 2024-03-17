export function useSessionKey() {
    const NourishUserList = 'nourish-user-list'
    const NourishUserInfoUpdated = 'nourish-user-info-updated'
    const NourishUserTotal = 'nourish-user-total'
    const NourishLogoutUserTotal = 'nourish-logout-user-total'
    const NourishLogoutUserList = 'nourish-logout-user-list'
    const NourishLogoutUserInfoUpdated = 'nourish-logout-user-info-updated'
    const NourishPromise = 'nourish-promise'

    return {
        NourishUserList, NourishUserInfoUpdated, NourishUserTotal,
        NourishLogoutUserTotal, NourishLogoutUserList, NourishLogoutUserInfoUpdated,
        NourishPromise
    }
}